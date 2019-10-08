package base

import (
	"github.com/elliotcourant/buffers"
)

type RowTupleHeader struct {
	TableId    uint64
	PrimaryKey [][]byte
}

func (r *RowTupleHeader) EncodeKey() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte('r')
	buf.AppendUint64(r.TableId)
	buf.AppendUint8(uint8(len(r.PrimaryKey)))
	for _, key := range r.PrimaryKey {
		buf.Append(key...)
	}
	return buf.Bytes()
}
