package nitrodb

import (
	"github.com/elliotcourant/buffers"
)

type Datum struct {
	TableId    uint8
	PrimaryKey [][]byte
	CommitTs   uint64
	Values     map[uint8][]byte
}

func (d Datum) GetColumnValue(columnId uint8) []byte {
	return d.Values[columnId]
}

func (d Datum) Encode() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendByte('d')
	buf.AppendUint8(d.TableId)
	buf.AppendUint8(uint8(len(d.PrimaryKey)))
	for _, key := range d.PrimaryKey {
		buf.Append(key...)
	}
	buf.AppendUint64(d.CommitTs)
	header, values := d.encodeValues()
	buf.AppendUint8(uint8(len(header)))
	for columnId, location := range header {
		buf.AppendUint8(columnId)
		buf.AppendUint16(location[0])
		buf.AppendUint16(location[1])
	}
	buf.Append(values...)
	return buf.Bytes()
}

func (d Datum) encodeValues() (map[uint8][2]uint16, []byte) {
	header := map[uint8][2]uint16{}
	buf := make([]byte, 0)
	for columnId, value := range d.Values {
		header[columnId] = [2]uint16{
			uint16(len(buf)),
			uint16(len(buf) + len(value)),
		}
		buf = append(buf, value...)
	}
	return header, buf
}

func (d *Datum) Decode(src []byte) {
	*d = Datum{
		Values: map[uint8][]byte{},
	}
	buf := buffers.NewBytesReader(src)
	buf.NextByte() // Get rid of the row prefix.
	d.TableId = buf.NextUint8()
	pKeySize := int(buf.NextUint8())
	d.PrimaryKey = make([][]byte, pKeySize)
	for i := 0; i < pKeySize; i++ {
		d.PrimaryKey[i] = buf.NextBytes()
	}
	d.CommitTs = buf.NextUint64()
	headerSize := int(buf.NextUint8())
	header := map[uint8][2]uint16{}
	for i := 0; i < headerSize; i++ {
		columnId := buf.NextUint8()
		start, end := buf.NextUint16(), buf.NextUint16()
		header[columnId] = [2]uint16{
			start,
			end,
		}
	}
	valueHeap := buf.NextBytes()
	for columnId, location := range header {
		d.Values[columnId] = valueHeap[location[0]:location[1]]
	}
}
