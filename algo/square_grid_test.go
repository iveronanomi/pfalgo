package algo

import (
	. "gopkg.in/check.v1"
)

type SquareGridSuite struct{}

var _ = Suite(&SquareGridSuite{})

func (s *SquareGridSuite) TestSquareGrid_Passable(c *C) {
	sg := NewSquareGrid(10, 5, LinearWalk)
	sg.AddWall(NewNode(2, 2, 0), 6, 1)

	passable := sg.Passable(NewNode(2, 2, 0))

	c.Assert(passable, Equals, false)
}

func (s *SquareGridSuite) TestSquareGrid_InBound_True(c *C) {
	sg := NewSquareGrid(10, 5, LinearWalk)

	inBound := sg.InBound(NewNode(1, 4, 0))

	c.Assert(inBound, Equals, true)
}

func (s *SquareGridSuite) TestSquareGrid_InBound_False(c *C) {
	sg := NewSquareGrid(10, 5, LinearWalk)

	inBound := sg.InBound(NewNode(0, 5, 0))

	c.Assert(inBound, Equals, false)
}

func (s *SquareGridSuite) TestSquareGrid_Neighbours_TopLeftCorrect(c *C) {
	sg := NewSquareGrid(5, 5, LinearWalk)

	neighbours := sg.Neighbours(NewNode(0, 0, 0))

	c.Assert(neighbours, HasLen, 2)
}

func (s *SquareGridSuite) TestSquareGrid_Neighbours_FourCorrect(c *C) {
	sg := NewSquareGrid(5, 5, LinearWalk)

	neighbours := sg.Neighbours(NewNode(1, 1, 0))

	c.Assert(neighbours, HasLen, 4)
}

func (s *SquareGridSuite) TestSquareGrid_Neighbours_WithWallCorrect(c *C) {
	sg := NewSquareGrid(5, 5, LinearWalk)
	sg.AddObstructions(NewNode(2, 1, 0))

	neighbours := sg.Neighbours(NewNode(2, 2, 0))

	c.Assert(neighbours, HasLen, 3)
}
