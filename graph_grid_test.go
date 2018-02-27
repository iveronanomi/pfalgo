package pfalgo

import (
	. "gopkg.in/check.v1"
)

type SquareGridSuite struct{}

var _ = Suite(&SquareGridSuite{})

func (s *SquareGridSuite) TestSquareGrid_Passable(c *C) {
	sg := NewSquareGrid(10, 5, LinearWalk, nil)
	sg.AddWall(2, 2, 6, 1)

	passable := sg.Passable(2, 2)

	c.Assert(passable, Equals, false)
}

func (s *SquareGridSuite) TestSquareGrid_InBound_True(c *C) {
	sg := NewSquareGrid(10, 5, LinearWalk, nil)

	inBound := sg.InBound(1, 4)

	c.Assert(inBound, Equals, true)
}

func (s *SquareGridSuite) TestSquareGrid_InBound_False(c *C) {
	sg := NewSquareGrid(10, 5, LinearWalk, nil)

	inBound := sg.InBound(0, 5)

	c.Assert(inBound, Equals, false)
}

func (s *SquareGridSuite) TestSquareGrid_Neighbours_TopLeftCorrect(c *C) {
	sg := NewSquareGrid(5, 5, LinearWalk, nil)

	neighbours := sg.Neighbours(Node{0, 0})

	c.Assert(neighbours, HasLen, 2)
}

func (s *SquareGridSuite) TestSquareGrid_Neighbours_FourCorrect(c *C) {
	sg := NewSquareGrid(5, 5, LinearWalk, nil)

	neighbours := sg.Neighbours(Node{1, 1})

	c.Assert(neighbours, HasLen, 4)
}

func (s *SquareGridSuite) TestSquareGrid_Neighbours_WithWallCorrect(c *C) {
	sg := NewSquareGrid(5, 5, LinearWalk, nil)
	sg.AddWall(2, 1, 1, 1)

	neighbours := sg.Neighbours(Node{2, 2})

	c.Assert(neighbours, HasLen, 3)
}
