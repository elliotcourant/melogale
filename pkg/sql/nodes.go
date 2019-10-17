package sql

import (
	"reflect"
)

var _ PlanStep = &CreateTablePlan{}
var _ PlanStep = &InsertTablePlan{}
var _ PlanStep = &ValuesListRenderer{}

type PlanStage []PlanStep

type PlanStep interface {
	Run(ctx ExecutionContext) error
	Explain() Explanation
}

type RowValue map[uint8]reflect.Value

type Receiver interface {
	ReceiveRow(row RowValue)
}
