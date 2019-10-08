package engine

var _ PlanNode = &AddColumnPlan{}
var _ PlanNode = &AddForeignKeyConstraintPlan{}
var _ PlanNode = &AddUniqueConstraintPlan{}
var _ PlanNode = &ColumnDoesExistPlan{}
var _ PlanNode = &CreateTablePlan{}
var _ PlanNode = &ObjectIdPlan{}
var _ PlanNode = &TableDoesExistPlan{}
var _ PlanNode = &TableDoesNotExistPlan{}
var _ PlanNode = &StatementNode{}
