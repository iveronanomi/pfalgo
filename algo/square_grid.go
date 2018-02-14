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

	visited map[Node]struct{} //list of visited nodes
	start   Node              //start node
	target  Node              //target node
}

// InBound is it node locatated in the grid
func (g *SquareGrid) InBound(node Node) bool {
	// log.Printf("%v |  0 <= %d < %d  && | 0 <= %d < %d", node, node[0], g.width, node[1], g.height)
	return (0 <= node[0] && node[0] < g.width) &&
		(0 <= node[1] && node[1] < g.height)
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
	nodes := []Node{
		{node[0] + 1, node[0]},
		{node[0], node[0] - 1},
		{node[0] - 1, node[0]},
		{node[0], node[0] + 1},
	}

	return NodeFilter(nodes, func(n Node) bool {
		return g.Passable(n) && g.InBound(n)
	})
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

// AddObstructions add list of walls as nodes
func (g *SquareGrid) AddObstructions(nodes ...Node) {
	if g.obstructions == nil {
		g.obstructions = make(map[Node]struct{})
	}

	for _, n := range nodes {
		if g.InBound(n) {
			g.obstructions[n] = struct{}{}
			// log.Printf("Add node: %v", n)
			continue
		}
		// log.Printf("NOTICE: not in bound: %v", n)
	}
}

// AddWall add wall according on Heigt, Width
func (g *SquareGrid) AddWall(start Node, height, width int) {
	// log.Printf("Draw a wall from %v width: %v, height: %v", start, height, width)
	for x := int(start[0]); x < width+int(start[0]); x++ {
		for y := int(start[1]); y < height+int(start[1]); y++ {
			g.AddObstructions(Node{uint32(x), uint32(y)})
		}
	}
}

// String create preview of grid
func (g *SquareGrid) String() string {
	out := ""
	for y := 0; y < int(g.height); y++ {
		for x := 0; x < int(g.width); x++ {
			node := Node{uint32(x), uint32(y)}
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
