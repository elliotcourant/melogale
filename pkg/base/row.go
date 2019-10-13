package base

import (
	"github.com/elliotcourant/buffers"
)

type Row struct {
	TableId    uint8
	PrimaryKey []Datum
	Datums     map[uint8]Datum
}

func (r Row) EncodeKey() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte(byte(RowPrefix))
	buf.AppendUint8(r.TableId)
	buf.AppendUint8(uint8(len(r.PrimaryKey)))
	for _, primaryKey := range r.PrimaryKey {
		buf.Append(primaryKey.Encode()...)
	}
	return buf.Bytes()
}

func (r Row) EncodeValue() []byte {
	buf := buffers.NewBytesBuffer()
	locations, values := r.encodeDatums()
	buf.AppendUint8(uint8(len(locations)))
	for columnId, offsets := range locations {
		buf.AppendUint8(columnId)
		buf.AppendUint16(offsets[0])
		buf.AppendUint16(offsets[1])
	}
	buf.Append(values...)
	return buf.Bytes()
}

func (r Row) encodeDatums() (map[uint8][2]uint16, []byte) {
	size := 0
	buf := buffers.NewBytesBuffer()
	locations := map[uint8][2]uint16{}
	for columnId, datum := range r.Datums {
		encoded := datum.Encode()
		buf.AppendRaw(encoded)
		locations[columnId] = [2]uint16{
			uint16(size),
			uint16(size + len(encoded)),
		}
		size += len(encoded)
	}
	return locations, buf.Bytes()
}

func (r *Row) DecodeKey(src []byte) {
	buf := buffers.NewBytesReader(src)
	buf.NextByte()
	r.TableId = buf.NextUint8()
	pkeySize := int(buf.NextUint8())
	r.PrimaryKey = make([]Datum, pkeySize)
	for i := 0; i < pkeySize; i++ {
		datum := Datum{}
		val := buf.NextBytes()
		datum.Decode(val)
		r.PrimaryKey[i] = datum
	}
}

func (r *Row) DecodeValue(src []byte) {
	buf := buffers.NewBytesReader(src)
	mapSize := buf.NextUint8()
	locations := map[uint8][2]uint16{}
	for i := uint8(0); i < mapSize; i++ {
		locations[buf.NextUint8()] = [2]uint16{
			buf.NextUint16(),
			buf.NextUint16(),
		}
	}
	r.Datums = map[uint8]Datum{}
	values := buf.NextBytes()
	for columnId, offset := range locations {
		datum := Datum{}
		val := values[offset[0]:offset[1]]
		datum.Decode(val)
		r.Datums[columnId] = datum
	}
}
