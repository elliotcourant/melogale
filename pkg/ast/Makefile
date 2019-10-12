.PHONY: generate

generate:
	stringer -type Context -output context.string.go context.go
	stringer -type ConstrType -output constr_type.string.go constr_type.go
	stringer -type ObjectType -output object_type.string.go object_type.go
	stringer -type SortByDir -output sort_by_dir.string.go sort_by_dir.go
	stringer -type StmtType -output stmt_type.string.go stmt_type.go
	stringer -type SubLinkType -output sub_link_type.string.go sub_link_type.go
	stringer -type SQLValueFunctionOp -output sql_value_function_op.string.go sql_value_function_op.go
	stringer -type TransactionStmtKind -output transaction_stmt_kind.string.go transaction_stmt_kind.go