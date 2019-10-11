package engine

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/base"
	"github.com/elliotcourant/melogale/pkg/store"
)

type PlanContext interface {
	GetTable(tableName string) (base.TableHeader, error)
	GetTableColumns(tableName string, columnNames ...string) ([]base.ColumnHeader, error)
	GetTableIndexes(tableName string, indexNames ...string) ([]base.IndexHeader, error)
}

type CreationContext interface {
	PlanContext

	Txn() store.Transaction
}

type ExecutionContext interface {
	PlanContext

	Txn() store.Transaction
}

type planContextBase struct {
	tables  map[string]base.TableHeader
	columns map[string]map[string]base.ColumnHeader
	indexes map[string]map[string]base.IndexHeader

	txn store.Transaction
}

func (p *planContextBase) GetTable(tableName string) (base.TableHeader, error) {
	if table, ok := p.tables[tableName]; ok {
		return table, nil
	}
	key := base.NewTableNamePrefix(tableName)
	value, err := p.txn.Get(key)
	if err != nil {
		return base.TableHeader{}, err
	}
	if len(value) == 0 {
		return base.TableHeader{}, fmt.Errorf("table [%s] does not exist", tableName)
	}
	table := base.TableHeader{}
	table.DecodeKey(key)
	table.DecodeValue(value)
	p.tables[tableName] = table
	return table, nil
}

func (p *planContextBase) GetTableColumns(tableName string, columnNames ...string) ([]base.ColumnHeader, error) {
	columns, ok := p.columns[tableName]
	if !ok {
		columns = map[string]base.ColumnHeader{}
		table, err := p.GetTable(tableName)
		if err != nil {
			return nil, err
		}
		itr := p.txn.Iterator()
		prefix := base.NewColumnNamePrefix(table.TableId, "")
		itr.Seek(prefix)
		for ; itr.ValidForPrefix(prefix); itr.Next() {
			k, v, err := itr.Value()
			if err != nil {
				return nil, err
			}
			column := base.ColumnHeader{}
			column.DecodeKey(k)
			column.DecodeValue(v)
			columns[column.Name] = column
		}
		p.columns[tableName] = columns
	}

	cols := make([]base.ColumnHeader, len(columnNames))
	if len(columnNames) == 0 {
		for _, col := range columns {
			cols = append(cols, col)
		}
	} else {
		for i, columnName := range columnNames {
			if col, ok := columns[columnName]; ok {
				cols[i] = col
			} else {
				return nil, fmt.Errorf("column [%s] for table [%s] does not exist", tableName, columnName)
			}
		}
	}

	return cols, nil
}

func (p *planContextBase) GetTableIndexes(tableName string, indexNames ...string) ([]base.IndexHeader, error) {
	panic("implement me")
}

func newPlanContext(txn store.Transaction) PlanContext {
	return &planContextBase{
		tables:  map[string]base.TableHeader{},
		columns: map[string]map[string]base.ColumnHeader{},
		indexes: map[string]map[string]base.IndexHeader{},
		txn:     txn,
	}
}
