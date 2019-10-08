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

func NewColumnNamePrefix(tableId uint64, columnName string) []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte('c')
	buf.AppendUint64(tableId)
	buf.AppendShortString(columnName)
	return buf.Bytes()
}
