package base

type TypeFamily uint8

const (
	IntFamily TypeFamily = iota + 1
	BoolFamily
	StringFamily
	TimeFamily
)
