package engine

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/parser"
	"testing"
	"time"
)

func TestPlannerBase_Plan(t *testing.T) {
	t.Run("create table", func(t *testing.T) {
		stmt, err := parser.Parse("CREATE TABLE users (id bigint primary key, account_id bigint references accounts (account_id), email text unique, password text, first_name text, last_name text);")
		if err != nil {
			panic(err)
		}
		p := NewPlanner()
		start := time.Now()
		result := p.PlanAll(stmt)
		fmt.Println("planning took:", time.Since(start))
		fmt.Println(result.Explain())
	})
}
