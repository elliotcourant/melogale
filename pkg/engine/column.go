package engine

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/base"
	"github.com/pingcap/parser/ast"
	"math"
)

type AddColumnPlan struct {
	tableName  string
	definition *ast.ColumnDef
}

func (p *plannerBase) NewAddColumnPlan(tableName string, def *ast.ColumnDef) PlanStack {
	stack := PlanStack{
		&AddColumnPlan{
			tableName:  tableName,
			definition: def,
		},
	}
	for _, option := range def.Options {
		switch option.Tp {
		case ast.ColumnOptionUniqKey:
			stack = append(stack, p.NewAddUniqueConstraintPlan(fmt.Sprintf("uq_%s_%s", tableName, def.Name.String())))
		case ast.ColumnOptionReference:
			stack = append(stack, p.NewAddForeignKeyConstraintPlan(tableName, []string{def.Name.String()}, option)...)
		}
	}
	return stack
}

func (a *AddColumnPlan) Explain() Explanation {
	return Explanation{
		Level:  2,
		Action: SET,
		Name:   "column header",
		Desc:   fmt.Sprintf("create column header: %s -> %s", a.tableName, a.definition.Name.String()),
		Key:    base.NewColumnNamePrefix(math.MaxUint64, a.definition.Name.String()),
	}
}

func (a *AddColumnPlan) Execute(ctx ExecuteContext) error {
	table, err := ctx.GetTable(a.tableName)
	if err != nil {
		return err
	}
	columnId, err := ctx.GetColumnId(a.tableName, a.definition.Name.String())
	if err != nil {
		return err
	}
	columnHeader := base.ColumnHeader{
		TableId:  table.TableId,
		ColumnId: columnId,
		Name:     a.definition.Name.String(),
	}
	return ctx.Txn().Set(columnHeader.EncodeKey(), columnHeader.EncodeValue())
}

type ColumnDoesExistPlan struct {
	tableName  string
	columnName string
}

func (c *ColumnDoesExistPlan) Explain() Explanation {
	return Explanation{
		Level:  1,
		Action: GET,
		Name:   "column header",
		Desc:   fmt.Sprintf("column with nane [%s] must exist on table [%s]", c.columnName, c.tableName),
		Key:    base.NewColumnNamePrefix(math.MaxUint64, c.columnName),
	}
}

func (c *ColumnDoesExistPlan) Execute(ctx ExecuteContext) error {
	panic("implement me")
}

func (p *plannerBase) NewColumnDoesExistPlan(tableName, columnName string) PlanNode {
	return &ColumnDoesExistPlan{
		tableName:  tableName,
		columnName: columnName,
	}
}
