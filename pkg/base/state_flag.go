package base

type State uint8

const (
	None      State = 0
	Writable  State = 1
	Readable  State = 1 << 1
	Deletable State = 1 << 2
)

func (i State) IsReadable() bool {
	return i&Readable != 0
}

func (i State) IsWritable() bool {
	return i&Writable != 0
}

func (i State) IsDeletable() bool {
	return i&Deletable != 0
}
