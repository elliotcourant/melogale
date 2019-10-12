package base

import (
	"github.com/elliotcourant/buffers"
)

type Column struct {
	ColumnId uint8
	Name     string
	Type     Type
	Flags    ColumnFlag
}

func (c Column) Encode() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendUint8(c.ColumnId)
	buf.AppendShortString(c.Name)
	buf.Append(c.Type.Encode()...)
	buf.AppendUint8(uint8(c.Flags))
	return buf.Bytes()
}

func (c *Column) Decode(src []byte) {
	buf := buffers.NewBytesReader(src)
	c.ColumnId = buf.NextUint8()
	c.Name = buf.NextShortString()
	c.Type.Decode(buf.NextBytes())
	c.Flags = ColumnFlag(buf.NextUint8())
}
