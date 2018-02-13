package algo

const (
	charWall   = " ⬛"
	charGrid   = " ⬜"
	charStart  = " Ⓢ"
	charFinish = " Ⓐ"
)

// Node ...
type Node [2]uint32

// SquareGrid ...
type SquareGrid struct {
	width  uint32
	height uint32
	walls  map[uint32]map[uint32]struct{}
}

// AddWalls add list of walls as nodes
func (g *SquareGrid) AddWalls(nodes ...Node) {
	for _, n := range nodes {
		g.walls[n[0]][n[1]] = struct{}{}
	}
}

// InBound is it node locatated in the grid
func (g *SquareGrid) InBound(node Node) bool {
	return (0 <= node[0] && node[0] < g.width) &&
		(0 <= node[1] && node[1] <= g.height)
}

// Passable is it passable node
func (g *SquareGrid) Passable(node Node) bool {
	_, ok := g.walls[node[0]][node[1]]
	return !ok
}

// Neighbours get all neibourhoods
func (g *SquareGrid) Neighbours(node Node) []Node {
	predictNodes := []Node{
		{node[0] + 1, node[0]},
		{node[0], node[0] - 1},
		{node[0] - 1, node[0]},
		{node[0], node[0] + 1},
	}

	result := make([]Node, 0, 4)
	for _, n := range predictNodes {
		if !g.Passable(n) || !g.InBound(n) {
			continue
		}
		result = append(result, n)
	}
	return result
}

// NewSquareGrid new instance of grid
func NewSquareGrid(width, height uint32, walls map[uint32]map[uint32]struct{}) *SquareGrid {
	return &SquareGrid{
		width:  width,
		height: height,
		walls:  walls,
	}
}

// String create preview of grid
func (g *SquareGrid) String() string {
	out := ""
	for i := 0; i <= int(g.height); i++ {
		for j := 0; j <= int(g.width); j++ {
			if _, ok := g.walls[uint32(i)][uint32(j)]; ok {
				out += charWall
				continue
			}
			out += charGrid
		}
		out += "\n"
	}
	return out
}
