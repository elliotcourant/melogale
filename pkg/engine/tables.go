package engine

var _ Table = tableBase{}

type Table interface {
	ID() uint8
	Name() string
	SetName(tableName string)
	Columns() []Column
	UpsertColumn(column Column)
	PrimaryKeys() []Column
}

type tableBase struct {
	id          uint8
	name        string
	columns     []Column
	primaryKeys []uint8
}

func (t tableBase) UpsertColumn(column Column) {
	panic("implement me")
}

func (t tableBase) SetName(tableName string) {
	panic("implement me")
}

func (t tableBase) ID() uint8 {
	panic("implement me")
}

func (t tableBase) Name() string {
	panic("implement me")
}

func (t tableBase) Columns() []Column {
	panic("implement me")
}

func (t tableBase) PrimaryKeys() []Column {
	panic("implement me")
}
