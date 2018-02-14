package algo

import (
	. "gopkg.in/check.v1"
)

type NodeSuite struct{}

var _ = Suite(&NodeSuite{})

func (s *NodeSuite) TestNodeIsPassable(c *C) {
	n := &AdvancedNode{}
	n.setFlag(flagNotPassable)

	r := n.IsPassable()

	c.Assert(r, Equals, false)
}

func (s *NodeSuite) TestNodeIsVisited(c *C) {
	n := &AdvancedNode{}
	n.setFlag(flagVisited)

	r := n.IsVisited()

	c.Assert(r, Equals, true)
}

func (s *NodeSuite) TestNodeIsTarget(c *C) {
	n := &AdvancedNode{}
	n.setFlag(flagTarget)

	r := n.IsTarget()

	c.Assert(r, Equals, true)
}

func (s *NodeSuite) TestNodeFlagsComplicated(c *C) {
	n := &AdvancedNode{}
	n.setFlag(flagNotPassable | flagStart)

	c.Assert(n.IsTarget(), Equals, false)
	c.Assert(n.IsPassable(), Equals, false)
	c.Assert(n.IsStart(), Equals, true)
	c.Assert(n.IsVisited(), Equals, false)
}

func (s *NodeSuite) TestNodeHasTwoFlags_ExpectYes(c *C) {
	n := &AdvancedNode{
		flags: (flagNotPassable | flagStart | flagTarget | flagVisited),
	}

	hf := n.hasFlag(flagNotPassable | flagStart)

	c.Assert(hf, Equals, true)
}

func (s *NodeSuite) TestNodeHasTwoFlags_ExpectNot(c *C) {
	n := &AdvancedNode{
		flags: (flagNotPassable | flagTarget | flagVisited),
	}

	hf := n.hasFlag(flagNotPassable | flagStart)

	c.Assert(hf, Equals, false)
}

func (s *NodeSuite) TestNodeSetPassable(c *C) {
	n := &AdvancedNode{
		flags: flagVisited,
	}
	n.removeFlag(flagVisited)

	visited := n.IsVisited()

	c.Assert(visited, Equals, false)
}
