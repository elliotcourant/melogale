package engine

import (
	"fmt"
	"github.com/elliotcourant/melogale/pkg/parser"
	"testing"
	"time"
)

func MustPlan(t *testing.T, query string) {
	stmt, err := parser.Parse(query)
	if err != nil {
		panic(err)
	}
	p := NewPlanner()
	start := time.Now()
	result := p.PlanAll(stmt)
	fmt.Println("Planning took:", time.Since(start))
	fmt.Println(result.Explain())
}

func TestPlannerBase_Plan(t *testing.T) {
	t.Run("create table", func(t *testing.T) {
		MustPlan(t, "CREATE TABLE users (id bigint primary key, account_id bigint references accounts (account_id), email text unique, password text, first_name text, last_name text);")
	})

	t.Run("insert", func(t *testing.T) {
		MustPlan(t, "INSERT INTO users (account_id, email, first_name, last_name) VALUES(123, 'me@me.com', 'Elliot', 'Courant');")
	})
}
