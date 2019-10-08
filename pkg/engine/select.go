package engine

import (
	"github.com/pingcap/parser/ast"
)

type SelectPlan struct {
	tableName string
	stmt      *ast.SelectStmt
}
