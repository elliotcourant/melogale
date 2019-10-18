package engine

type Table interface {
	ID() uint8
	Name() string
	Columns() []Column
	PrimaryKeys() []Column
}
