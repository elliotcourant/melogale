package base

import (
	"github.com/elliotcourant/buffers"
)

type TableHeader struct {
	TableId uint64
	Name    string
}

func (t *TableHeader) EncodeKey() []byte {
	return NewTableNamePrefix(t.Name)
}

func (t *TableHeader) EncodeValue() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendUint64(t.TableId)
	return buf.Bytes()
}

func (t *TableHeader) DecodeKey(src []byte) {
	buf := buffers.NewBytesReader(src)
	buf.NextByte()
	t.Name = buf.NextShortString()
}

func (t *TableHeader) DecodeValue(src []byte) {
	buf := buffers.NewBytesReader(src)
	t.TableId = buf.NextUint64()
}

func NewTableNamePrefix(tableName string) []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte('t')
	buf.AppendShortString(tableName)
	return buf.Bytes()
}
