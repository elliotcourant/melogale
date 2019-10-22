package engine

var _ Table = &tableBase{}

type Table interface {
	ID() uint8
	Name() string
	SetName(tableName string)

	Columns() []Column
	UpsertColumn(column Column)
	GetColumn(columnId uint8) Column
	GetColumnByName(columnName string) Column

	PrimaryKeys() []Column

	Indexes() []Index
	UpsertIndex(index Index)
	GetIndex(indexId uint8) Index
	GetIndexByName(indexName string) Index
}

type tableBase struct {
	id          uint8
	name        string
	columns     []Column
	primaryKeys []uint8
}

func (t *tableBase) GetColumnByName(columnName string) Column {
	panic("implement me")
}

func (t *tableBase) Indexes() []Index {
	panic("implement me")
}

func (t *tableBase) UpsertIndex(index Index) {
	panic("implement me")
}

func (t *tableBase) GetIndex(indexId uint8) Index {
	panic("implement me")
}

func (t *tableBase) GetIndexByName(indexName string) Index {
	panic("implement me")
}

func (t *tableBase) GetColumn(columnId uint8) Column {
	panic("implement me")
}

func (t *tableBase) UpsertColumn(column Column) {
	panic("implement me")
}

func (t *tableBase) SetName(tableName string) {
	panic("implement me")
}

func (t *tableBase) ID() uint8 {
	panic("implement me")
}

func (t *tableBase) Name() string {
	panic("implement me")
}

func (t *tableBase) Columns() []Column {
	panic("implement me")
}

func (t *tableBase) PrimaryKeys() []Column {
	panic("implement me")
}
