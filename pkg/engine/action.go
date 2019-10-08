package engine

//go:generate stringer --type=Action --output=action.string.go action.go
type Action int

const (
	NONE Action = iota
	SCAN
	GET
	SET
	ID
)
