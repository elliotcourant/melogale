package base

import (
	"github.com/elliotcourant/buffers"
)

type Index struct {
	IndexId uint8
	Name    string
	Columns map[uint8]uint8
}

func (i Index) Encode() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendUint8(i.IndexId)
	buf.AppendShortString(i.Name)
	buf.AppendUint8(uint8(len(i.Columns)))
	for columnId, position := range i.Columns {
		buf.AppendUint8(columnId)
		buf.AppendUint8(position)
	}
	return buf.Bytes()
}

func (i *Index) Decode(src []byte) {
	buf := buffers.NewBytesReader(src)
	i.IndexId = buf.NextUint8()
	i.Name = buf.NextShortString()
	i.Columns = map[uint8]uint8{}
	colSize := int(buf.NextUint8())
	for x := 0; x < colSize; x++ {
		i.Columns[buf.NextUint8()] = buf.NextUint8()
	}
}
