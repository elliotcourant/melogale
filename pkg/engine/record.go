package engine

type record struct {
	primaryKeyColumns []uint8
	columnValues      map[uint8][]byte
}

func (r record) PrimaryKey() [][]byte {
	keys := make([][]byte, len(r.primaryKeyColumns))
	for i, pkey := range r.primaryKeyColumns {
		keys[i] = r.columnValues[pkey]
	}
	return keys
}

func (r record) ColumnIds() []uint8 {
	columnIds := make([]uint8, 0)
	for columnId := range r.columnValues {
		columnIds = append(columnIds, columnId)
	}
	return columnIds
}

func (r record) GetColumn(columnId uint8) ([]byte, bool) {
	value, ok := r.columnValues[columnId]
	return value, ok
}

func newRecord() Record {
	return record{
		primaryKeyColumns: make([]uint8, 0),
		columnValues:      map[uint8][]byte{},
	}
}
