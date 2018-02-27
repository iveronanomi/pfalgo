package pfalgo

import (
	. "gopkg.in/check.v1"
)

type PriorityQueueSuite struct{}

var _ = Suite(&PriorityQueueSuite{})

func (s *PriorityQueueSuite) TestPriorityPush(c *C) {
	q := NewPriorityQueue()
	q.Add(Node{0, 1}, 2)
	q.Add(Node{0, 2}, 3)
	q.Add(Node{0, 3}, 3)
	q.Add(Node{0, 4}, 4)
	q.Add(Node{0, 5}, 2)

	c.Assert(q.Get(), Equals, Node{0, 1})
	c.Assert(q.Get(), Equals, Node{0, 5})
	c.Assert(q.Get(), Equals, Node{0, 3})
	c.Assert(q.Get(), Equals, Node{0, 2})
	c.Assert(q.Get(), Equals, Node{0, 4})
}
