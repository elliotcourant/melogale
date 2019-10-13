package sql

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/ast"
	"github.com/elliotcourant/melogale/pkg/base"
)

func (p *planner) CreateTable(stmt ast.CreateStmt) (PlanStage, error) {
	tableName := *stmt.Relation.Relname
	_, ok, err := p.GetTable(tableName)
	if err != nil {
		return nil, err
	}
	if ok {
		return nil, fmt.Errorf("a table with name [%s] already exists", tableName)
	}
	return PlanStage{
		CreateTablePlan{
			stmt:      stmt,
			tableName: tableName,
		},
	}, nil
}

type CreateTablePlan struct {
	stmt      ast.CreateStmt
	tableName string
	key       []byte
	value     []byte
}

func (c CreateTablePlan) Run(ctx ExecutionContext) error {
	table := base.Table{
		TableId: 0,
		Name:    c.tableName,
		Columns: map[string]base.Column{},
		Indexes: map[string]base.Index{},
	}
	columnCount := uint8(0)
	for _, elt := range c.stmt.TableElts.Items {
		switch item := elt.(type) {
		case ast.ColumnDef:
			columnCount++
			column := base.Column{
				ColumnId: columnCount,
				Name:     *item.Colname,
				Type:     base.Type{},
				Flags:    0,
			}

			for _, c := range item.Constraints.Items {
				switch constraint := c.(type) {
				case ast.Constraint:
					switch constraint.Contype {
					case ast.CONSTR_PRIMARY:
						column.Flags |= base.ColumnPrimaryKey
					}
				}
			}

			column.Type = base.GetType(*item.TypeName)

			table.Columns[column.Name] = column
		default:
			panic(fmt.Sprintf("cannot handle elt [%T]", item))
		}
	}

	return ctx.Set(table.EncodeKey(), table.EncodeValue())
}

func (c CreateTablePlan) Explain() Explanation {
	return Explanation{
		Order:       0,
		Action:      0,
		Name:        "create table",
		Description: fmt.Sprintf("create table: %s", c.tableName),
		Key:         c.key,
		Value:       c.value,
		Cost:        1,
	}
}
