package engine

import (
	"github.com/elliotcourant/melogale/pkg/base"
)

type Record interface {
	PrimaryKey() [][]byte
	ColumnIds() []uint8
	GetColumn(columnId uint8) ([]byte, bool)
}

type TableReader interface {
	Table() base.Table
	Seek(primaryKey ...[]byte)
	Valid() bool
	Next() bool
	CurrentPrimaryKey() [][]byte
	Record() Record
	Get(primaryKey ...[]byte) (Record, bool, error)
}

type SchemaReader interface {
	Seek(tableName []byte)
	Valid() bool
	Next() bool
	CurrentTableKey() []byte
	Table() base.Table
	Get(tableName []byte) (base.Table, bool, error)
}

type IndexReader interface {
	Table() base.Table
	Columns() []base.Column
	Index() base.Index
	Seek(columnValue ...[]byte)
	CurrentIndexKey() [][]byte
	Valid() bool
	Next() bool
	Item() Record
	Get(columnValue ...[]byte) (Record, bool, error)
}

type UniqueConstraintReader interface {
	Table() base.Table
	Columns() []base.Column
	UniqueConstraint() int

	IsUnique(columnValue ...[]byte) (unique bool, err error)
}

type ForeignConstraintReader interface {
	LocalTable() base.Table
	ForeignTable() base.Table

	LocalColumns() []base.Column
	ForeignColumns() []base.Column

	ForeignKey() int

	IsValid(columnValue ...[]byte) (valid bool, err error)
}
