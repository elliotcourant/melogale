package engine

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/base"
	"github.com/elliotcourant/melogale/pkg/store"
)

var _ ExecuteContext = &execContext{}

type ExecuteContext interface {
	SetObjectId(objectType base.ObjectType, name string, id uint64)
	GetObjectId(objectType base.ObjectType, name string) (uint64, error)

	AddTable(table base.TableHeader)
	GetTable(name string) (base.TableHeader, error)

	AddColumnId(table, column string, id uint8)
	GetColumnId(table, column string) (uint8, error)

	AddIndex(table string, index interface{})
	GetIndexes(table string) []interface{}

	Txn() store.Transaction
}

func NewExecutionContext(transaction store.Transaction) ExecuteContext {
	return &execContext{
		objIds:    map[base.ObjectType]map[string]uint64{},
		columnIds: map[string]map[string]uint8{},
		tables:    map[string]base.TableHeader{},
	}
}

type execContext struct {
	txn store.Transaction

	objIds    map[base.ObjectType]map[string]uint64
	columnIds map[string]map[string]uint8

	tables map[string]base.TableHeader
}

func (e *execContext) AddIndex(table string, index interface{}) {
	panic("implement me")
}

func (e *execContext) GetIndexes(table string) []interface{} {
	panic("implement me")
}

func (e *execContext) SetObjectId(objectType base.ObjectType, name string, id uint64) {
	e.objIds[objectType][name] = id
}

func (e *execContext) GetObjectId(objectType base.ObjectType, name string) (uint64, error) {
	names, ok := e.objIds[objectType]
	if !ok {
		return 0, fmt.Errorf("could not resolve object Ids for [%s]", objectType)
	}
	id, ok := names[name]
	if !ok {
		return 0, fmt.Errorf("could not resolve object Id for [%s - %s]", objectType, name)
	}
	return id, nil
}

func (e *execContext) AddTable(table base.TableHeader) {
	e.tables[table.Name] = table
}

func (e *execContext) GetTable(name string) (base.TableHeader, error) {
	table, ok := e.tables[name]
	if !ok {
		return table, fmt.Errorf("could not resolve table [%s]", name)
	}
	return table, nil
}

func (e *execContext) AddColumnId(table, column string, id uint8) {
	e.columnIds[table][column] = id
}

func (e *execContext) GetColumnId(table, column string) (uint8, error) {
	columns, ok := e.columnIds[table]
	if !ok {
		return 0, fmt.Errorf("could not resolve columns for table [%s]", table)
	}
	id, ok := columns[column]
	if !ok {
		return 0, fmt.Errorf("could not resolve column [%s] for table [%s]", column, table)
	}
	return id, nil
}

func (e *execContext) Txn() store.Transaction {
	return e.txn
}
