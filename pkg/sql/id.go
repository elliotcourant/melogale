package sql

import (
	"math"
)

type UniqueColumn uint16

func (u UniqueColumn) TableId() uint8 {
	return uint8((u >> 8) & math.MaxUint8)
}

func (u UniqueColumn) ColumnId() uint8 {
	return uint8(u & math.MaxUint8)
}

func NewUniqueColumn(tableId, columnId uint8) UniqueColumn {
	b := uint16(tableId)
	b = b << 8
	b |= uint16(columnId)
	return UniqueColumn(b)
}
