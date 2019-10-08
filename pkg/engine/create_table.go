package engine

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/base"
	"github.com/pingcap/parser/ast"
)

type CreateTablePlan struct {
	tableName string
	columns   []*ast.ColumnDef
}

func (c *CreateTablePlan) FailurePlan() PlanStack {
	panic("implement me")
}

func (c *CreateTablePlan) Name() string {
	return fmt.Sprintf("CreateTablePlan_%s", c.tableName)
}

func (c *CreateTablePlan) Explain() Explanation {
	return Explanation{
		Level:  2,
		Action: SET,
		Name:   "table header",
		Desc:   fmt.Sprintf("create table header: %s", c.tableName),
		Key:    base.NewTableNamePrefix(c.tableName),
	}
}

func (c *CreateTablePlan) Execute(ctx ExecuteContext) error {
	// Add the column Ids for this table
	for i, col := range c.columns {
		ctx.AddColumnId(c.tableName, col.Name.String(), uint8(i)+1)
	}

	tableId, err := ctx.GetObjectId(base.TABLE, c.tableName)
	if err != nil {
		return err
	}
	header := base.TableHeader{
		TableId: tableId,
		Name:    c.tableName,
	}
	return ctx.Txn().Set(header.EncodeKey(), header.EncodeValue())
}

func (p *plannerBase) CreateTable(stmt *ast.CreateTableStmt) PlanStack {
	tableName := stmt.Table.Name.String()
	plan := []PlanNode{
		p.NewTableDoesNotExistPlan(tableName),
		p.NewObjectIdPlan(base.TABLE, tableName),
		&CreateTablePlan{
			tableName: tableName,
			columns:   stmt.Cols,
		},
	}

	for _, col := range stmt.Cols {
		plan = append(plan, p.NewAddColumnPlan(tableName, col)...)
	}

	return plan
}
