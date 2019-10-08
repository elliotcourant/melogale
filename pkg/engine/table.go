package engine

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/base"
)

type TableDoesNotExistPlan struct {
	tableName string
	prefix    []byte
}

func (p *plannerBase) NewTableDoesNotExistPlan(tableName string) PlanNode {
	return &TableDoesNotExistPlan{
		tableName: tableName,
		prefix:    base.NewTableNamePrefix(tableName),
	}
}

func (t *TableDoesNotExistPlan) Explain() Explanation {
	return Explanation{
		Level:  0,
		Action: GET,
		Name:   "table header",
		Desc:   fmt.Sprintf("table with name [%s] must not exist", t.tableName),
		Key:    t.prefix,
	}
}

func (t *TableDoesNotExistPlan) Execute(ctx ExecuteContext) error {
	header, err := ctx.Txn().Get(t.prefix)
	if err != nil {
		return fmt.Errorf("could not verify table with name [%s] does not exist: %v", t.tableName, err)
	}
	if header == nil {
		return fmt.Errorf("a table with name [%s] already exists", t.tableName)
	}
	return nil
}

type TableDoesExistPlan struct {
	tableName string
	prefix    []byte
}

func (p *plannerBase) NewTableDoesExistPlan(tableName string) PlanNode {
	return &TableDoesExistPlan{
		tableName: tableName,
		prefix:    base.NewTableNamePrefix(tableName),
	}
}

func (t *TableDoesExistPlan) Explain() Explanation {
	return Explanation{
		Level:  0,
		Action: GET,
		Name:   "table header",
		Desc:   fmt.Sprintf("table with name [%s] must exist", t.tableName),
		Key:    t.prefix,
	}
}

func (t *TableDoesExistPlan) Execute(ctx ExecuteContext) error {
	header, err := ctx.Txn().Get(t.prefix)
	if err != nil {
		return fmt.Errorf("could not verify table with name [%s] does exist: %v", t.tableName, err)
	}
	if header == nil {
		return fmt.Errorf("a table with name [%s] does not exist", t.tableName)
	}
	table := &base.TableHeader{}
	table.DecodeKey(t.prefix)
	table.DecodeValue(header)
	ctx.AddTable(*table)
	return nil
}
