package engine

import (
	"encoding/hex"
	"fmt"
	"github.com/pingcap/parser/ast"
	"strings"
)

var _ Planner = &plannerBase{}

type Explanation struct {
	Level  int
	Action Action
	Name   string
	Desc   string
	Key    []byte
}

func (e Explanation) String() string {
	return strings.TrimSpace(fmt.Sprintf("[%02d] %-5s %-25s %-80s %-10s", e.Level, e.Action, e.Name, e.Desc, hex.EncodeToString(e.Key)))
}

type PlanStack []PlanNode

func (p PlanStack) Explain() string {
	s := strings.TrimSpace(fmt.Sprintf("%-4s %-5s %-25s %-80s %-10s", "LVL", "ACTN", "NAME", "DESC", "KEY"))
	s += "\n"
	s += fmt.Sprintf("=====================================================================================================\n")
	for i, e := range p {
		s += fmt.Sprintf("%s", e.Explain())
		if i != len(p)-1 {
			s += "\n"
		}
	}
	return s
}

func (p PlanStack) Execute(ctx ExecuteContext) error {
	for _, e := range p {
		if err := e.Execute(ctx); err != nil {
			return err
		}
	}
	return nil
}

type PlanNode interface {
	Explain() Explanation
	Execute(ctx ExecuteContext) error
}

type Planner interface {
	Plan(tree ast.StmtNode) PlanStack
	PlanAll(trees []ast.StmtNode) PlanStack
}

func NewPlanner() Planner {
	return &plannerBase{}
}

type plannerBase struct {
}

func (p *plannerBase) PlanAll(trees []ast.StmtNode) PlanStack {
	stack := make(PlanStack, 0)
	for _, tree := range trees {
		stack = append(stack, p.Plan(tree)...)
	}
	return stack
}

func (p *plannerBase) Plan(tree ast.StmtNode) PlanStack {
	stack := func() PlanStack {
		switch stmt := tree.(type) {
		case *ast.CreateTableStmt:
			return p.CreateTable(stmt)
		default:
			panic("not implemented")
		}
	}()
	return append(PlanStack{
		StatementNode{
			stmt: tree.Text(),
		},
	}, stack...)
}

type StatementNode struct {
	stmt string
}

func (s StatementNode) Explain() Explanation {
	return Explanation{
		Level:  -1,
		Action: NONE,
		Name:   "query",
		Desc:   s.stmt,
		Key:    nil,
	}
}

func (s StatementNode) Execute(ctx ExecuteContext) error {
	return nil
}
