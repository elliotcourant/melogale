package engine

import (
	"fmt"
)

type GetAllTableIndexesPlan struct {
	tableName string
}

func (g GetAllTableIndexesPlan) Explain() Explanation {
	return Explanation{
		Level:  2,
		Action: SCAN,
		Name:   "index header",
		Desc:   fmt.Sprintf("get all indexes for table: %s", g.tableName),
		Key:    nil,
	}
}

func (g GetAllTableIndexesPlan) Execute(ctx ExecuteContext) error {
	panic("implement me")
}

func (g GetAllTableIndexesPlan) Name() string {
	panic("implement me")
}

func (g GetAllTableIndexesPlan) AlternatePlan() PlanStack {
	panic("implement me")
}

func (p *plannerBase) NewGetAllTableIndexesPlan(tableName string) PlanNode {
	return &GetAllTableIndexesPlan{tableName: tableName}
}

type TableDoesHaveOptimalIndex struct {
	tableName string
	columns   []string
}

func (t TableDoesHaveOptimalIndex) Explain() Explanation {
	panic("implement me")
}

func (t TableDoesHaveOptimalIndex) Execute(ctx ExecuteContext) error {
	panic("implement me")
}

func (t TableDoesHaveOptimalIndex) Name() string {
	panic("implement me")
}

func (t TableDoesHaveOptimalIndex) AlternatePlan() PlanStack {
	panic("implement me")
}
