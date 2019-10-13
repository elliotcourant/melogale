package sql

var _ PlanStep = &CreateTablePlan{}
var _ PlanStep = &InsertTablePlan{}

type PlanStage []PlanStep

type PlanStep interface {
	Run(ctx ExecutionContext) error
	Explain() Explanation
}
