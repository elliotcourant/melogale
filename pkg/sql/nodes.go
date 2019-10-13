package sql

var _ PlanStep = &CreateTablePlan{}
var _ PlanStep = &InsertTablePlan{}
var _ PlanStep = &ValuesListPlan{}

type PlanStage []PlanStep

type PlanStep interface {
	Run(ctx ExecutionContext) error
	Explain() Explanation
}
