package base

import (
	"github.com/elliotcourant/buffers"
)

type ForeignKeyHeader struct {
	TableId             uint64
	ReferencedTableId   uint64
	ForeignKeyId        uint8
	ColumnIds           map[uint8]interface{}
	ReferencedColumnIds map[uint8]interface{}
	Name                string
}

func NewForeignKeyPrefix(tableId uint64, name string) []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte('f')
	buf.AppendUint64(tableId)
	buf.AppendShortString(name)
	return buf.Bytes()
}
