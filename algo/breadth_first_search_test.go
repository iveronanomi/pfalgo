package algo

import (
	. "gopkg.in/check.v1"
)

type BreadthFirstSearchSuite struct{}

var _ = Suite(&BreadthFirstSearchSuite{})

func (s *BreadthFirstSearchSuite) TestBreadthFirstSearch1(c *C) {
	g := SimpleGraph{
		Nodes: map[string][]string{
			"A": {"B"},
			"B": {"A", "C", "D"},
			"C": {"A"},
			"D": {"A", "E"},
			"E": {"B"},
		},
	}

	BreadthFirstSearch1(g, "A")
}
