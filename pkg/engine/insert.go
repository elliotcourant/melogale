package engine

import (
	"fmt"
	"github.com/pingcap/parser/ast"
	"strings"
)

type InsertPlan struct {
	tableName string
	columns   []string
	values    [][]ast.ExprNode
}

func (i InsertPlan) Explain() Explanation {
	return Explanation{
		Level:  3,
		Action: SET,
		Name:   "insert",
		Desc:   fmt.Sprintf("insert %d record(s) into: %s -> %s", len(i.values), i.tableName, strings.Join(i.columns, ", ")),
		Key:    nil,
	}
}

func (i InsertPlan) Execute(ctx ExecuteContext) error {
	// indexes := ctx.GetIndexes(i.tableName)
	panic("implement me")
}

func (i InsertPlan) Name() string {
	panic("implement me")
}

func (i InsertPlan) AlternatePlan() PlanStack {
	panic("implement me")
}

func (p *plannerBase) Insert(stmt *ast.InsertStmt) PlanStack {
	stack := make(PlanStack, 0)
	tableName := stmt.Table.TableRefs.Left.(*ast.TableSource).Source.(*ast.TableName).Name.String()
	stack = append(stack, p.NewTableDoesExistPlan(tableName))
	stack = append(stack, p.NewGetAllTableColumnsPlan(tableName))
	columnNames := make([]string, 0)
	for _, column := range stmt.Columns {
		columnNames = append(columnNames, column.Name.String())
		stack = append(stack, p.NewColumnDoesExistPlan(tableName, column.Name.String()))
	}

	stack = append(stack, p.NewGetAllTableIndexesPlan(tableName))
	stack = append(stack, p.NewGetAllTableUniqueConstraints(tableName))

	stack = append(stack, &InsertPlan{
		tableName: tableName,
		columns:   columnNames,
		values:    stmt.Lists,
	})

	return stack
}

type InsertIndexDatumsPlan struct {
	tableName string
}
