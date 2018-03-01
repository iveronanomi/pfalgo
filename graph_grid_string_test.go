package pfalgo

import (
	. "gopkg.in/check.v1"
)

type SquareGridStringSuite struct{}

var _ = Suite(&SquareGridStringSuite{})

func (s *SquareGridStringSuite) TestSquareGridString_Default(c *C) {
	sg := NewSquareGrid(5, 5, LinearWalk, nil)
	sg.AddWall(2, 0, 5, 1)
	sg.SetStart(0, 0)
	sg.SetTarget(4, 4)
	sg.Visit(1, 1)
	expectedString := " Ⓐ ◻ ◼ ◻ ◻\n\r ◻ ◆ ◼ ◻ ◻\n\r ◻ ◻ ◼ ◻ ◻\n\r ◻ ◻ ◼ ◻ ◻\n\r ◻ ◻ ◼ ◻ Ⓩ\n\r"

	out := String(sg, nil)

	c.Assert(out, Equals, expectedString)
}

func (s *SquareGridStringSuite) TestSquareGridString_CustomSymbols(c *C) {
	sg := NewSquareGrid(5, 5, LinearWalk, nil)
	sg.AddWall(2, 0, 5, 1)
	sg.SetStart(0, 0)
	sg.SetTarget(4, 4)
	sg.Visit(1, 1)
	customSymbols := map[GridGraphSymbols]string{
		SymbolObstruction: "1",
		SymbolGrid:        "0",
		SymbolStart:       "A",
		SymbolTarget:      "Z",
		SymbolVisited:     "+",
	}
	expectedString := "A0100\n\r0+100\n\r00100\n\r00100\n\r0010Z\n\r"

	out := String(sg, customSymbols)

	c.Assert(out, Equals, expectedString)
}
