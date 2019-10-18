package engine

import (
	"github.com/elliotcourant/buffers"
)

var _ TableReader = &tableReader{}

const (
	rowPrefix = 'r'
)

func NewTableReader(txn Transaction, table Table) TableReader {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte(rowPrefix)
	buf.AppendUint8(table.ID())
	return &tableReader{
		rowPrefix: buf.Bytes(),
		table:     table,
		txn:       txn,
		itr:       txn.Iterator(),
	}
}

type tableReader struct {
	rowPrefix []byte
	table     Table
	txn       Transaction
	itr       Iterator
}

func (t tableReader) prefix() []byte {
	p := make([]byte, len(t.rowPrefix))
	copy(p, t.rowPrefix)
	return p
}

func (t tableReader) Table() Table {
	return t.table
}

func (t tableReader) Seek(primaryKeys ...[]byte) {
	t.itr.Seek(t.buildKey(primaryKeys...))
}

func (t tableReader) Valid() bool {
	return t.itr.Valid()
}

func (t tableReader) Next() bool {
	t.itr.Next()
	return t.Valid()
}

func (t tableReader) CurrentPrimaryKey() [][]byte {
	numKeys := len(t.table.PrimaryKeys())
	key, _, _ := t.itr.Item()
	// Throw out the row prefix and the table id.
	buf := buffers.NewBytesReader(key[2:])
	keys := make([][]byte, numKeys)
	for i := 0; i < numKeys; i++ {
		keys[i] = buf.NextBytes()
	}
	return keys
}

func (t tableReader) Record() Record {
	_, _, err := t.itr.Item()
	if err != nil {
		panic(err)
	}
	return nil
}

func (t tableReader) Get(primaryKeys ...[]byte) (Record, bool, error) {
	value, err := t.txn.Get(t.buildKey(primaryKeys...))
	if err != nil || len(value) == 0 {
		return nil, false, err
	}
	panic("implement me")
}

func (t tableReader) buildKey(primaryKeys ...[]byte) []byte {
	prefix := t.prefix()
	if len(primaryKeys) == 0 {
		return prefix
	}
	buf := buffers.NewBytesBuffer()
	buf.AppendRaw(prefix)
	for _, primaryKey := range primaryKeys {
		buf.Append(primaryKey...)
	}
	return buf.Bytes()
}
