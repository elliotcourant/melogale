package sql

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/ast"
	"github.com/elliotcourant/melogale/pkg/base"
	"sort"
)

func (p *planner) Insert(stmt ast.InsertStmt) (PlanStage, error) {
	tableName := *stmt.Relation.Relname
	table, ok, err := p.GetTable(tableName)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("a table with name [%s] does not exist", tableName)
	}

	stage := make(PlanStage, 0)

	columns := make([]base.Column, len(stmt.Cols.Items))

	if len(stmt.Cols.Items) > 0 {
		for i, col := range stmt.Cols.Items {
			switch resTarget := col.(type) {
			case ast.ResTarget:
				columnName := *resTarget.Name
				column, ok := table.Columns[columnName]
				if !ok {
					return nil, fmt.Errorf("a column with name [%s] does not exist for table [%s]", columnName, tableName)
				}
				columns[i] = column
			default:
				panic("invalid column")
			}
		}
	} else {
		for _, column := range table.Columns {
			columns = append(columns, column)
		}
		sort.Slice(columns, func(i, j int) bool {
			return columns[i].ColumnId < columns[j].ColumnId
		})
	}

	switch v := stmt.SelectStmt.(type) {
	case ast.SelectStmt:
		if v.ValuesLists == nil || len(v.ValuesLists) == 0 {
			panic("insert from select query not implemented")
		} else {
			stage = append(stage, ValuesListPlan{
				stmt:    v,
				table:   table,
				columns: columns,
			})
		}
	default:
		panic(fmt.Sprintf("insert from [%T] not implemented", v))
	}

	stage = append(stage, InsertTablePlan{
		stmt:    stmt,
		table:   table,
		columns: columns,
	})

	if len(table.Indexes) > 0 {
		panic("add insert index plans")
	}

	return stage, nil
}

type InsertTablePlan struct {
	stmt    ast.InsertStmt
	table   base.Table
	columns []base.Column
}

func (i InsertTablePlan) Run(ctx ExecutionContext) error {
	rows := make([]base.Row, 0)
	rowValues := ctx.GetValues()
	for _, rowValue := range rowValues {
		row := base.Row{
			TableId:    i.table.TableId,
			PrimaryKey: make([]base.Datum, 0),
			Datums:     map[uint8]base.Datum{},
		}
		for columnName, value := range rowValue {
			col := i.table.Columns[columnName]
			if col.Flags.IsPrimaryKey() {
				row.PrimaryKey = append(row.PrimaryKey, value)
			}
			row.Datums[col.ColumnId] = value
		}
		rows = append(rows, row)
	}
	for _, row := range rows {
		if err := ctx.Set(row.EncodeKey(), row.EncodeValue()); err != nil {
			return err
		}
	}
	return nil
}

func (i InsertTablePlan) Explain() Explanation {
	panic("implement me")
}

func (p *planner) ValuesList(values [][]ast.Node) error {
	panic("test")
}

type ValuesListPlan struct {
	stmt    ast.SelectStmt
	table   base.Table
	columns []base.Column
}

func (v ValuesListPlan) Run(ctx ExecutionContext) error {
	valueList := v.stmt.ValuesLists
	for _, rowValues := range valueList {
		row := RowValue{}
		for c, cell := range rowValues {
			col := v.columns[c]
			datum := base.Datum{
				Type:  col.Type.Family,
				Value: nil,
			}
			for {
				switch val := cell.(type) {
				case ast.A_Const:
					cell = val.Val
					continue
				case ast.String:
					datum.Value = val.Str
				case ast.Integer:
					datum.Value = val.Ival
				}
				break
			}
			row[col.Name] = datum
		}
		ctx.StoreValue(row)
	}
	return nil
}

func (v ValuesListPlan) Explain() Explanation {
	panic("implement me")
}
