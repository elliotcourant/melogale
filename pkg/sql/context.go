package sql

import (
	"github.com/dgraph-io/badger"
	"github.com/elliotcourant/melogale/pkg/base"
	"github.com/elliotcourant/melogale/pkg/engine"
)

type RowValue map[string]base.Datum

var _ AssistanceContext = &assistanceContext{}

type AssistanceContext interface {
	GetTable(tableName string) (base.Table, bool, error)
	GetValues() []RowValue
	StoreValue(rowValue RowValue)
	ClearValues()
}

type ExecutionContext interface {
	engine.Transaction
	AssistanceContext
}

type executionContext struct {
	engine.Transaction
	AssistanceContext
}

func newExecutionContext(txn engine.Transaction) ExecutionContext {
	return newExecutionContextEx(txn, newAssistanceContext(txn))
}

func newExecutionContextEx(txn engine.Transaction, ctx AssistanceContext) ExecutionContext {
	return &executionContext{
		Transaction:       txn,
		AssistanceContext: ctx,
	}
}

type PlanningContext interface {
	AssistanceContext
}

func newAssistanceContext(txn engine.Transaction) AssistanceContext {
	return &assistanceContext{
		txn:    txn,
		tables: map[string]base.Table{},
		values: make([]RowValue, 0),
	}
}

type assistanceContext struct {
	txn    engine.Transaction
	tables map[string]base.Table
	values []RowValue
}

func (a *assistanceContext) ClearValues() {
	a.values = make([]RowValue, 0)
}

func (a *assistanceContext) GetValues() []RowValue {
	return a.values
}

func (a *assistanceContext) StoreValue(rowValue RowValue) {
	a.values = append(a.values, rowValue)
}

func (a *assistanceContext) GetTable(tableName string) (base.Table, bool, error) {
	table, ok := a.tables[tableName]
	if !ok {
		key := base.NewTableNameKey(tableName)
		value, err := a.txn.Get(key)
		if err == badger.ErrKeyNotFound || len(value) == 0 {
			return table, false, nil
		} else if err != nil {
			return table, false, err
		}
		table.DecodeValue(value)
		a.tables[tableName] = table
	}
	return table, true, nil
}
