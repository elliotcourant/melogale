package engine

type SchemaWriter interface {
	CreateTable(table Table) error
}

type TableWriter interface {
	Table() Table
	Insert(records ...Record) error
	Update(records ...Record) error
	Delete(records ...Record) error
}

type IndexWriter interface {
	Insert(records ...Record) error
	Swap(old, new Record) error
	Delete(records ...Record) error
}
