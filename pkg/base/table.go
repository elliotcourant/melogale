package base

import (
	"github.com/elliotcourant/buffers"
)

type Table struct {
	TableId uint8
	Name    string
	Columns map[string]Column
	Indexes map[string]Index
}

func (t Table) EncodeValue() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendUint8(t.TableId)
	buf.AppendShortString(t.Name)
	buf.AppendUint8(uint8(len(t.Columns)))
	for _, column := range t.Columns {
		buf.Append(column.Encode()...)
	}
	buf.AppendUint8(uint8(len(t.Indexes)))
	for _, index := range t.Indexes {
		buf.Append(index.Encode()...)
	}
	return buf.Bytes()
}

func (t *Table) DecodeValue(src []byte) {
	buf := buffers.NewBytesReader(src)
	t.TableId = buf.NextUint8()
	t.Name = buf.NextShortString()
	t.Columns = map[string]Column{}
	t.Indexes = map[string]Index{}
	colSize := int(buf.NextUint8())
	for x := 0; x < colSize; x++ {
		col := Column{}
		col.Decode(buf.NextBytes())
		t.Columns[col.Name] = col
	}
	indexSize := int(buf.NextUint8())
	for x := 0; x < indexSize; x++ {
		ind := Index{}
		ind.Decode(buf.NextBytes())
		t.Indexes[ind.Name] = ind
	}
}

func (t Table) EncodeKey() []byte {
	return NewTableNameKey(t.Name)
}

func NewTableNameKey(tableName string) []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte(byte(TablePrefix))
	if len(tableName) > 0 {
		buf.AppendShortString(tableName)
	}
	return buf.Bytes()
}
