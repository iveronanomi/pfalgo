package algo

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type BreadthFirstSearchSuite struct{}

var _ = Suite(&BreadthFirstSearchSuite{})

func (s *BreadthFirstSearchSuite) BreadthFirstSearch(c *C) {
	g := SimpleGraph{
		Nodes: map[string][]string{
			"A": {"B"},
			"B": {"A", "C", "D"},
			"C": {"A"},
			"D": {"A", "E"},
			"E": {"B"},
		},
	}

	BreadthFirstSearch(g, "A")

	c.Assert(true, Equals, false)
}
