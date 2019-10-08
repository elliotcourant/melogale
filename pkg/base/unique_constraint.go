package base

import (
	"github.com/elliotcourant/buffers"
)

type UniqueConstraintHeader struct {
	TableId            uint64
	UniqueConstraintId uint8
	Name               string
	ColumnIds          map[uint8]interface{}
}

func (u *UniqueConstraintHeader) EncodeKey() []byte {
	return NewUniqueConstraintPrefix(u.TableId, u.Name)
}

func (u *UniqueConstraintHeader) EncodeValue() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendUint8(u.UniqueConstraintId)
	buf.AppendUint8(uint8(len(u.ColumnIds)))
	for columnId := range u.ColumnIds {
		buf.AppendUint8(columnId)
	}
	return buf.Bytes()
}

func NewUniqueConstraintPrefix(tableId uint64, name string) []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte('x')
	buf.AppendUint64(tableId)
	buf.AppendShortString(name)
	return buf.Bytes()
}
