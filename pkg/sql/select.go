package sql

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/ast"
	"github.com/elliotcourant/melogale/pkg/base"
)

func (p *planner) Select(stmt ast.SelectStmt) PlanStage {
	return nil
}

type SelectRenderer struct {
	stmt ast.SelectStmt
}

func (p *planner) ValuesListRenderer(stmt ast.SelectStmt, columns []base.Column, receiver Receiver) (PlanStage, error) {
	return PlanStage{
		&ValuesListRenderer{
			stmt:     stmt,
			columns:  columns,
			receiver: receiver,
		},
	}, nil
}

type ValuesListRenderer struct {
	stmt     ast.SelectStmt
	columns  []base.Column
	receiver Receiver
}

func (v ValuesListRenderer) Run(ctx ExecutionContext) error {
	// for _, rowItems := range v.stmt.ValuesLists {
	// 	// rowValue := RowValue{}
	// 	// for _, value := range rowItems {
	// 	// 	for {
	// 	// 		switch v := value.(type) {
	// 	// 		case ast.A_Const:
	// 	// 			value
	// 	// 		}
	// 	// 	}
	// 	// }
	// }
	return nil
}

func (v ValuesListRenderer) Explain() Explanation {
	return Explanation{
		Action:      0,
		Name:        "values",
		Description: fmt.Sprintf("prepare %d row(s) for %T", len(v.stmt.ValuesLists), v.receiver),
		Key:         nil,
		Value:       nil,
		Cost:        0,
	}
}
