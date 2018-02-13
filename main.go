package main

import (
	"github.com/iveronanomi/pfalgo/algo"
)

func main() {
	g := algo.SimpleGraph{
		Nodes: map[string][]string{
			"A": {"B"},
			"B": {"A", "C", "D"},
			"C": {"A"},
			"D": {"A", "E"},
			"E": {"B", "F"},
			"F": {"A"},
		},
	}

	algo.BreadthFirstSearch(g, "A")
}
