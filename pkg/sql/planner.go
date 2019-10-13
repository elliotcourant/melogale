package sql

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/ast"
	"github.com/elliotcourant/melogale/pkg/engine"
)

var _ Planner = &planner{}
var _ QueryPlan = &queryPlan{}

type Planner interface {
	Build(node ast.SyntaxTree) (QueryPlan, error)
}

type QueryPlan interface {
	Run() error
}

type queryPlan struct {
	ex         ExecutionContext
	statements []PlanStage
}

func (q *queryPlan) Run() error {
	for _, statement := range q.statements {
		for _, step := range statement {
			if err := step.Run(q.ex); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewPlanner(txn engine.Transaction) Planner {
	return &planner{
		txn:               txn,
		AssistanceContext: newAssistanceContext(txn),
	}
}

type planner struct {
	txn engine.Transaction
	AssistanceContext
}

func (p *planner) Build(node ast.SyntaxTree) (QueryPlan, error) {
	plan := &queryPlan{
		ex:         newExecutionContextEx(p.txn, p.AssistanceContext),
		statements: make([]PlanStage, 0),
	}

	for _, stmtStep := range node.Statements {

		for {
			switch stmtStep.(type) {
			case ast.RawStmt:
				stmtStep = stmtStep.(ast.RawStmt).Stmt
			default:
				goto BUILD
			}
		}
	BUILD:

		stage, err := make(PlanStage, 0), error(nil)
		switch stmt := stmtStep.(type) {
		case ast.CreateStmt:
			stage, err = p.CreateTable(stmt)
		case ast.InsertStmt:
			stage, err = p.Insert(stmt)
		default:
			panic(fmt.Sprintf("%T statements are not supported at this time", stmt))
		}

		if err != nil {
			return nil, err
		}

		plan.statements = append(plan.statements, stage)
	}

	return plan, nil
}
