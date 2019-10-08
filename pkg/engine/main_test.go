package engine

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
)

func CreateGraph(stack PlanStack) {
	graphAst, _ := gographviz.ParseString(`digraph G {}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}

	levels := make([]PlanStack, len(stack))
	for _, node := range stack {
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

	output := graph.String()
	fmt.Println(output)
}
