package engine

type SequenceWriter interface {
	NewTableID() (uint8, error)
}

type SchemaWriter interface {
	CreateTable(table Table) error
}

type TableWriter interface {
	Table() Table
	Upsert(records ...Record) error
	Delete(records ...Record) error
}

type IndexWriter interface {
	Insert(records ...Record) error
	Swap(old, new Record) error
	Delete(records ...Record) error
}
