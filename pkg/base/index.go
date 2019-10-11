package base

import (
	"github.com/elliotcourant/buffers"
)

type IndexHeader struct {
	TableId uint64
	IndexId uint8
	Name    string
	Columns map[uint8]uint8
}

func (i IndexHeader) EncodeKey() []byte {
	return NewIndexNamePrefix(i.TableId, i.Name)
}

func (i IndexHeader) EncodeValue() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendUint8(i.IndexId)
	buf.AppendUint8(uint8(len(i.Columns)))
	for columnId, index := range i.Columns {
		buf.AppendUint8(columnId)
		buf.AppendUint8(index)
	}
	return buf.Bytes()
}

func NewIndexNamePrefix(tableId uint64, indexName string) []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte('i')
	buf.AppendUint64(tableId)
	if len(indexName) > 0 {
		buf.AppendShortString(indexName)
	}
	return buf.Bytes()
}
