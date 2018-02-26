package algo

import "image/color"

// INode interface
type INode interface {
	Position() (uint32, uint32)
	Equal(node INode) bool
}

// IDraw interface
type IDraw interface {
	Point(x int, y int, pointType color.RGBA)
	Visit(x int, y int)
	Save()
}
