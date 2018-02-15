package algo

import (
	. "gopkg.in/check.v1"
)

type NodeSuite struct{}

var _ = Suite(&NodeSuite{})

func (s *NodeSuite) TestNodeIsPassable(c *C) {
	n := NewNode(0, 0, 0)
	n.setFlag(statusNotPassable)

	r := n.IsPassable()

	c.Assert(r, Equals, false)
}

func (s *NodeSuite) TestNodeIsVisited(c *C) {
	n := NewNode(0, 0, 0)
	n.setFlag(statusVisited)

	r := n.IsVisited()

	c.Assert(r, Equals, true)
}

func (s *NodeSuite) TestNodeIsTarget(c *C) {
	n := NewNode(0, 0, 0)
	n.setFlag(statusTarget)

	r := n.IsTarget()

	c.Assert(r, Equals, true)
}

func (s *NodeSuite) TestNodeFlagsComplicated(c *C) {
	n := NewNode(0, 0, 0)
	n.setFlag(statusNotPassable | statusStart)

	c.Assert(n.IsTarget(), Equals, false)
	c.Assert(n.IsPassable(), Equals, false)
	c.Assert(n.IsStart(), Equals, true)
	c.Assert(n.IsVisited(), Equals, false)
}

func (s *NodeSuite) TestNodeHasTwoFlags_ExpectYes(c *C) {
	n := &Node{
		status: (statusNotPassable | statusStart | statusTarget | statusVisited),
	}

	hf := n.hasFlag(statusNotPassable | statusStart)

	c.Assert(hf, Equals, true)
}

func (s *NodeSuite) TestNodeHasTwoFlags_ExpectNot(c *C) {
	n := &Node{
		status: (statusNotPassable | statusTarget | statusVisited),
	}

	hf := n.hasFlag(statusNotPassable | statusStart)

	c.Assert(hf, Equals, false)
}

func (s *NodeSuite) TestNodeSetPassable(c *C) {
	n := &Node{
		status: statusVisited,
	}
	n.removeFlag(statusVisited)

	visited := n.IsVisited()

	c.Assert(visited, Equals, false)
}
