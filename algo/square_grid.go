package algo

import (
	"log"
)

const (
	charObstruction = " ◼"
	charGrid        = " ◻"
	charStart       = " Ⓐ"
	charTarget      = " Ⓩ"
	charVisited     = " ○"
)

// SquareGrid space
type SquareGrid struct {
	width        uint32            //width of grid
	height       uint32            //height of grid
	obstructions map[Node]struct{} //list of obstructions in the grid, not passable

	weights map[Node]map[Node]int //movements cost to node
	// todo (weights): potential problem of duplicating and lost elements for map[{0,0}]map[{0,1}] = 1; map[{0,1}]map[{0,0}] = 1,
	// the same movements and cost
	visited map[Node]struct{} //list of visited nodes
	start   Node              //start node
	target  Node              //target node
}

// InBound is it node locatated in the grid
func (g *SquareGrid) InBound(node Node) bool {
	x, y := node.Position()
	return (0 <= x && x < g.width) && (0 <= y && y < g.height)
}

// Passable is it passable node
func (g *SquareGrid) Passable(node Node) bool {
	_, ok := g.obstructions[node]
	return !ok
}

// NodeFilter function
func NodeFilter(nodes []Node, filter func(n Node) bool) []Node {
	result := []Node{}
	for _, n := range nodes {
		if filter(n) {
			result = append(result, n)
		}
	}
	return result
}

// Neighbours get all neibourhoods
func (g *SquareGrid) Neighbours(node Node) []Node {
	x, y := node.Position()
	nodes := []Node{
		NewNode(x+1, y, 0),
		NewNode(x, y-1, 0),
		NewNode(x-1, y, 0),
		NewNode(x, y+1, 0),
	}

	return NodeFilter(nodes, func(n Node) bool {
		return g.Passable(n) && g.InBound(n)
	})
}

// Cost of movements from `node` to `node`
func (g *SquareGrid) Cost(from, to Node) int {
	return 0 //todo
}

// NewSquareGrid new instance of grid
func NewSquareGrid(width, height uint32) *SquareGrid {
	log.Printf("Create rectangle grid, height: %v, width: %v", height, width)
	return &SquareGrid{
		width:        width,
		height:       height,
		obstructions: map[Node]struct{}{},
		weights:      map[Node]map[Node]int{},
	}
}

// AddObstructions add list of walls as nodes
func (g *SquareGrid) AddObstructions(nodes ...Node) {
	for _, n := range nodes {
		if g.InBound(n) {
			g.obstructions[n] = struct{}{}
			// log.Printf("Add obstruction node: %v", n)
			continue
		}
		// log.Printf("NOTICE: not in bound: %v", n)
	}
}

// AddWall add wall according on Heigt, Width
func (g *SquareGrid) AddWall(start Node, height, width int) {
	xPos, yPos := start.Position()
	log.Printf("Draw a wall %v x %v from %v point", height, width, start)
	for x := int(xPos); x < width+int(xPos); x++ {
		for y := int(yPos); y < height+int(yPos); y++ {
			g.AddObstructions(NewNode(uint32(x), uint32(y), 0))
		}
	}
}

// String create preview of grid
func (g *SquareGrid) String() string {
	out := ""
	for y := 0; y < int(g.height); y++ {
		for x := 0; x < int(g.width); x++ {
			node := NewNode(uint32(x), uint32(y), 0)
			if g.target == node {
				out += charTarget
				continue
			}
			if g.start == node {
				out += charStart
				continue
			}
			if !g.Passable(node) {
				out += charObstruction
				continue
			}
			if _, ok := g.visited[node]; ok {
				out += charVisited
				continue
			}
			out += charGrid
		}
		out += "\n\r"
	}
	return out
}

// Image of curent grid state
func (g *SquareGrid) Image() {
	// todo; use it for snapshot of current grid state, for creating a gif in future
}

// Visit set node of grid as visited (only for)
func (g *SquareGrid) Visit(node Node) {
	if g.visited == nil {
		g.visited = make(map[Node]struct{})
	}
	g.visited[node] = struct{}{}
}

// Start set start node (only for draw)
func (g *SquareGrid) Start(node Node) {
	g.start = node
}

// Target set start node (only for draw)
func (g *SquareGrid) Target(node Node) {
	g.target = node
}
