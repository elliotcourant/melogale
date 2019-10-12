package base

import (
	"github.com/elliotcourant/buffers"
)

type Type struct {
	Family TypeFamily
	Size   uint8
}

func (t Type) Encode() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendUint8(uint8(t.Family))
	buf.AppendUint8(t.Size)
	return buf.Bytes()
}

func (t *Type) Decode(src []byte) {
	buf := buffers.NewBytesReader(src)
	t.Family = TypeFamily(buf.NextUint8())
	t.Size = buf.NextUint8()
}
