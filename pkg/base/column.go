package base

import (
	"github.com/elliotcourant/buffers"
)

type ColumnHeader struct {
	TableId  uint64
	ColumnId uint8
	Name     string
}

func (c *ColumnHeader) EncodeKey() []byte {
	return NewColumnNamePrefix(c.TableId, c.Name)
}

func (c *ColumnHeader) EncodeValue() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendUint8(c.ColumnId)
	return buf.Bytes()
}

func (c *ColumnHeader) DecodeKey(src []byte) {
	buf := buffers.NewBytesReader(src)
	buf.NextByte() // Skip prefix byte
	c.TableId = buf.NextUint64()
	c.Name = buf.NextShortString()
}

func (c *ColumnHeader) DecodeValue(src []byte) {
	buf := buffers.NewBytesReader(src)
	c.ColumnId = buf.NextUint8()
}

func NewColumnNamePrefix(tableId uint64, columnName string) []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte('c')
	buf.AppendUint64(tableId)
	if len(columnName) > 0 {
		buf.AppendShortString(columnName)
	}
	return buf.Bytes()
}
