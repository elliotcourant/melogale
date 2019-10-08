package engine

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/base"
	"github.com/pingcap/parser/ast"
	"math"
	"strings"
)

type AddForeignKeyConstraintPlan struct {
	name string
}

func (a AddForeignKeyConstraintPlan) FailurePlan() PlanStack {
	panic("implement me")
}

func (a AddForeignKeyConstraintPlan) Name() string {
	return fmt.Sprintf("AddForeignKeyConstraintPlan_%s", a.name)
}

func (a AddForeignKeyConstraintPlan) Explain() Explanation {
	return Explanation{
		Level:  4,
		Action: SET,
		Name:   "foreign key constraint",
		Desc:   fmt.Sprintf("create foreign key constraint header: %s", a.name),
		Key:    base.NewForeignKeyPrefix(math.MaxUint64, a.name),
	}
}

func (a AddForeignKeyConstraintPlan) Execute(ctx ExecuteContext) error {
	panic("implement me")
}

func (p *plannerBase) NewAddForeignKeyConstraintPlan(table string, columns []string, option *ast.ColumnOption) PlanStack {
	stack := make(PlanStack, 0)
	referenceTable := option.Refer.Table.Name.String()
	stack = append(stack, p.NewTableDoesExistPlan(referenceTable))

	referenceColumns := make([]string, 0)
	for _, col := range option.Refer.IndexColNames {
		referenceColumns = append(referenceColumns, col.Column.Name.String())
		stack = append(stack, p.NewColumnDoesExistPlan(referenceTable, col.Column.Name.String()))
	}
	stack = append(stack, &AddForeignKeyConstraintPlan{
		name: fmt.Sprintf(
			"fk_%s_%s_%s_%s",
			table,
			strings.Join(columns, "_"),
			referenceTable,
			strings.Join(referenceColumns, "_"),
		),
	})
	return stack
}
