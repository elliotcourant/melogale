package engine

type Column interface {
	Name() string
	Type() Type
}
