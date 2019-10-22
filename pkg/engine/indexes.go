package engine

var _ Index = &indexBase{}

type Index interface {
	ID() uint8
	Name() string
	SetName(indexName string)
	Columns() []Column
}

type indexBase struct {
	id      uint8
	name    string
	table   Table
	columns []uint8
}

func (i *indexBase) ID() uint8 {
	return i.id
}

func (i *indexBase) Name() string {
	return i.name
}

func (i *indexBase) SetName(indexName string) {
	i.name = indexName
}

func (i *indexBase) Columns() []Column {
	cols := make([]Column, len(i.columns))
	for x, columnId := range i.columns {
		cols[x] = i.table.GetColumn(columnId)
	}
	return cols
}
