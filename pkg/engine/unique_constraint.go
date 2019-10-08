package engine

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/base"
	"math"
)

type AddUniqueConstraintPlan struct {
	name string
}

func (a AddUniqueConstraintPlan) Explain() Explanation {
	return Explanation{
		Level:  1,
		Action: SET,
		Name:   "unique constraint",
		Desc:   fmt.Sprintf("create unique constraint header: %s", a.name),
		Key:    base.NewUniqueConstraintPrefix(math.MaxUint64, a.name),
	}
}

func (a AddUniqueConstraintPlan) Execute(ctx ExecuteContext) error {
	panic("implement me")
}

func (p *plannerBase) NewAddUniqueConstraintPlan(name string) PlanNode {
	return &AddUniqueConstraintPlan{name: name}
}
