package sql

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/ast"
	"github.com/elliotcourant/melogale/pkg/base"
)

func (p *planner) Insert(stmt ast.InsertStmt) (PlanStage, error) {
	tableName := *stmt.Relation.Relname
	table, ok, err := p.GetTable(tableName)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("a table with name [%s] does not exist", tableName)
	}

	stage := make(PlanStage, 0)

	hasValuesClause := true
	switch v := stmt.SelectStmt.(type) {
	case ast.SelectStmt:
		if v.ValuesLists == nil || len(v.ValuesLists) == 0 {
			hasValuesClause = false
			panic("insert from select query not implemented")
		}
	default:
		panic(fmt.Sprintf("insert from [%T] not implemented", v))
	}

	stage = append(stage, InsertTablePlan{
		stmt:      stmt,
		tableName: tableName,
		hasValues: hasValuesClause,
	})

	if len(table.Indexes) > 0 {
		panic("add insert index plans")
	}

	return stage, nil
}

type InsertTablePlan struct {
	stmt      ast.InsertStmt
	tableName string
	hasValues bool
}

func (i InsertTablePlan) Run(ctx ExecutionContext) error {
	rows := make([]base.Row, 0)
	if !i.hasValues {

	} else {
		valueList := i.stmt.SelectStmt.(ast.SelectStmt).ValuesLists

	}

	for _, row := range rows {
		if err := ctx.Set(row.EncodeKey(), row.EncodeValue()); err != nil {
			return err
		}
	}
	return nil
}

func (i InsertTablePlan) Explain() Explanation {
	panic("implement me")
}
