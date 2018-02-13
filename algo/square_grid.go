package algo

import (
	"log"
)

const (
	charObstruction = " ⬛"
	charGrid        = " ⬜"
	charStart       = " Ⓢ"
	charFinish      = " Ⓐ"
)

// Node of square grid
type Node [2]uint32

// SquareGrid space
type SquareGrid struct {
	width        uint32            //width of grid
	height       uint32            //height of grid
	obstructions map[Node]struct{} //list of obstructions in the grid, not passable
}

// AddObstructions add list of walls as nodes
func (g *SquareGrid) AddObstructions(nodes ...Node) {
	if g.obstructions == nil {
		g.obstructions = make(map[Node]struct{})
	}

	for _, n := range nodes {
		if g.InBound(n) {
			g.obstructions[n] = struct{}{}
			log.Printf("Add node: %v", n)
			continue
		}
		log.Printf("NOTICE: not in bound: %v", n)
	}
}

// AddWall add wall according on Heigt, Width
func (g *SquareGrid) AddWall(start Node, height, width int) {
	log.Printf("Draw a wall from %v width: %v, height: %v", start, height, width)
	for x := int(start[0]); x < width+int(start[0]); x++ {
		for y := int(start[1]); y < height+int(start[1]); y++ {
			g.AddObstructions(Node{uint32(x), uint32(y)})
		}
	}
}

// InBound is it node locatated in the grid
func (g *SquareGrid) InBound(node Node) bool {
	log.Printf("%v |  0 <= %d < %d  && | 0 <= %d < %d", node, node[0], g.width, node[1], g.height)
	return (0 <= node[0] && node[0] < g.width) &&
		(0 <= node[1] && node[1] < g.height)
}

// Passable is it passable node
func (g *SquareGrid) Passable(node Node) bool {
	_, ok := g.obstructions[node]
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
func NewSquareGrid(width, height uint32, obstructions map[Node]struct{}) *SquareGrid {
	log.Printf("Create rectangle grid, width: %v, height: %v, obstructions: %v", width, height, obstructions)
	return &SquareGrid{
		width:        width,
		height:       height,
		obstructions: obstructions,
	}
}

// String create preview of grid
func (g *SquareGrid) String() string {
	out := ""
	for x := 0; x < int(g.height); x++ {
		for y := 0; y < int(g.width); y++ {
			if !g.Passable(Node{uint32(x), uint32(y)}) {
				out += charObstruction
				continue
			}
			out += charGrid
		}
		log.Println(x)
		out += "\n\r"
	}
	return out
}
