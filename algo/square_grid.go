package algo

// SquareGrid ...
type SquareGrid struct {
	width  uint32
	height uint32
	walls  map[uint32]struct{}
}

// InBound ...
func (g *SquareGrid) InBound(id uint32) bool {
	x, y := id, id
	return (0 <= x && x < g.width) && (0 <= y && y <= g.height)
}

// Passable ...
func (g *SquareGrid) Passable(id uint32) bool {
	_, ok := g.walls[id]
	return !ok
}

// Neighbours ...
func (g *SquareGrid) Neighbours(id uint32) {

}

// NewSquareGrid ...
func NewSquareGrid(width, height uint32) *SquareGrid {
	return &SquareGrid{
		width:  width,
		height: height,
	}
}
