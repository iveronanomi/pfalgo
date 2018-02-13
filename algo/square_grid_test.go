package algo

import (
	. "gopkg.in/check.v1"
)

type SquareGridSuite struct{}

var _ = Suite(&SquareGridSuite{})

func (s *SquareGridSuite) TestSquareGrid_Passable(c *C) {
	walls := map[uint32]map[uint32]struct{}{
		2: {
			2: struct{}{},
			3: struct{}{},
			4: struct{}{},
			5: struct{}{},
			6: struct{}{},
			7: struct{}{},
		},
	}
	sg := NewSquareGrid(10, 5, walls)

	passable := sg.Passable(Node{2, 2})

	c.Assert(passable, Equals, false)
}

func (s *SquareGridSuite) TestSquareGrid_InBound_True(c *C) {
	sg := NewSquareGrid(10, 5, nil)

	inBound := sg.InBound(Node{1, 5})

	c.Assert(inBound, Equals, true)
}

func (s *SquareGridSuite) TestSquareGrid_InBound_False(c *C) {
	sg := NewSquareGrid(10, 5, nil)

	inBound := sg.InBound(Node{0, 5})

	c.Assert(inBound, Equals, false)
}

func (s *SquareGridSuite) TestSquareGrid_Neighbours_TopLeftCorrect(c *C) {
	sg := NewSquareGrid(5, 5, nil)

	neighbours := sg.Neighbours(Node{0, 0})

	c.Assert(neighbours, HasLen, 2)
}

func (s *SquareGridSuite) TestSquareGrid_Neighbours_FourCorrect(c *C) {
	sg := NewSquareGrid(5, 5, nil)

	neighbours := sg.Neighbours(Node{1, 1})

	c.Assert(neighbours, HasLen, 4)
}

func (s *SquareGridSuite) TestSquareGrid_Neighbours_WithWallCorrect(c *C) {
	walls := map[uint32]map[uint32]struct{}{
		2: {
			1: struct{}{},
		},
	}
	sg := NewSquareGrid(5, 5, walls)

	neighbours := sg.Neighbours(Node{2, 2})

	c.Assert(neighbours, HasLen, 3)
}
