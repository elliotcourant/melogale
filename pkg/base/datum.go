package base

type Datum []byte

func (d Datum) Encode() []byte {
	return d
}

func (d *Datum) Decode(src []byte) {
	*d = src
}
