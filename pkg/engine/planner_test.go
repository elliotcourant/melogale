package engine

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
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

		graphAst, _ := gographviz.ParseString(`digraph G {}`)
		graph := gographviz.NewGraph()
		if err := gographviz.Analyse(graphAst, graph); err != nil {
			panic(err)
		}

		levels := make([]PlanStack, len(result))
		for _, node := range result {
			levels[node.Explain().Level+1] = append(levels[node.Explain().Level+1], node)
		}

		sortedPlan := make(PlanStack, 0)
		for _, level := range levels {
			for _, item := range level {
				sortedPlan = append(sortedPlan, item)
			}
		}
		fmt.Println(sortedPlan.Explain())

		previous := ""
		for _, level := range levels {
			for _, item := range level {
				graph.AddNode(graph.Name, item.Name(), map[string]string{})
				if len(previous) > 0 {
					graph.AddEdge(previous, item.Name(), true, map[string]string{})
				}
				previous = item.Name()
			}
		}
		//
		// graph.
		//
		// graph.AddNode("G", "a", nil)
		// graph.AddNode("G", "b", nil)
		// graph.AddEdge("a", "b", true, nil)
		output := graph.String()
		fmt.Println(output)
	})
}
