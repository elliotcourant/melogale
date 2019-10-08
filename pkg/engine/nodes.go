package engine

var _ PlanNode = &AddColumnPlan{}
var _ PlanNode = &AddForeignKeyConstraintPlan{}
var _ PlanNode = &AddUniqueConstraintPlan{}
var _ PlanNode = &ColumnDoesExistPlan{}
var _ PlanNode = &CreateTablePlan{}
var _ PlanNode = &GetAllTableColumnsPlan{}
var _ PlanNode = &GetAllTableIndexesPlan{}
var _ PlanNode = &GetAllTableUniqueConstraints{}
var _ PlanNode = &InsertPlan{}
var _ PlanNode = &ObjectIdPlan{}
var _ PlanNode = &TableDoesExistPlan{}
var _ PlanNode = &TableDoesNotExistPlan{}
var _ PlanNode = &TableDoesHaveOptimalIndex{}
var _ PlanNode = &StatementNode{}
