package engine

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/base"
	"math"
)

type AddUniqueConstraintPlan struct {
	name string
}

func (a AddUniqueConstraintPlan) AlternatePlan() PlanStack {
	panic("implement me")
}

func (a AddUniqueConstraintPlan) Name() string {
	return fmt.Sprintf("AddUniqueConstraintPlan_%s", a.name)
}

func (a AddUniqueConstraintPlan) Explain() Explanation {
	return Explanation{
		Level:  4,
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

type GetAllTableUniqueConstraints struct {
	tableName string
}

func (g GetAllTableUniqueConstraints) Explain() Explanation {
	return Explanation{
		Level:  2,
		Action: SCAN,
		Name:   "unique constraint",
		Desc:   fmt.Sprintf("get all unique constraints for table: %s", g.tableName),
		Key:    base.NewUniqueConstraintPrefix(math.MaxUint64, ""),
	}
}

func (g GetAllTableUniqueConstraints) Execute(ctx ExecuteContext) error {
	panic("implement me")
}

func (g GetAllTableUniqueConstraints) Name() string {
	panic("implement me")
}

func (g GetAllTableUniqueConstraints) AlternatePlan() PlanStack {
	panic("implement me")
}

func (p *plannerBase) NewGetAllTableUniqueConstraints(tableName string) PlanNode {
	return &GetAllTableUniqueConstraints{tableName: tableName}
}
