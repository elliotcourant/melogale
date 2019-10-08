package base

import (
	"github.com/elliotcourant/buffers"
)

func NewObjectIdPrefix(objectType ObjectType, name string) []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte('s')
	buf.AppendInt16(int16(objectType))
	buf.AppendShortString(name)
	return buf.Bytes()
}
