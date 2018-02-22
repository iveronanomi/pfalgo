package algo

import (
	. "gopkg.in/check.v1"
)

type PriorityQueueSuite struct{}

var _ = Suite(&PriorityQueueSuite{})

func (s *PriorityQueueSuite) TestPriorityPush(c *C) {
	q := NewPriorityQueue()
	q.Add(NewNode(0, 1, 0), 2)
	q.Add(NewNode(0, 2, 0), 3)
	q.Add(NewNode(0, 3, 0), 3)
	q.Add(NewNode(0, 4, 0), 4)
	q.Add(NewNode(0, 5, 0), 2)

	c.Assert(q.Get(), Equals, NewNode(0, 1, 0))
	c.Assert(q.Get(), Equals, NewNode(0, 5, 0))
	c.Assert(q.Get(), Equals, NewNode(0, 3, 0))
	c.Assert(q.Get(), Equals, NewNode(0, 2, 0))
	c.Assert(q.Get(), Equals, NewNode(0, 4, 0))
}
