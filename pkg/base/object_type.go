package base

//go:generate stringer --type=ObjectType --output=object_type.string.go object_type.go
type ObjectType int

const (
	TABLE ObjectType = iota
	COLUMN
)
