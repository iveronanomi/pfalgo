package algo

import (
	"log"
	"image/color"
)

const (
	charObstruction = " ◼"
	charGrid        = " ◻"
	charStart       = " Ⓐ"
	charTarget      = " Ⓩ"
	charVisited     = " ◆"
)

// WalkBehaviour ..
type WalkBehaviour uint8

const (
	// DiagonalWalk ...
	DiagonalWalk WalkBehaviour = 1 << iota
	// LinearWalk ...
	LinearWalk
	// AllWalk - walks all directions Diagonaly and Linear
	AllWalk
)

// SquareGrid space
type SquareGrid struct {
	width         uint32        //width of grid
	height        uint32        //height of grid
	walkBehaviour WalkBehaviour //is it possible to walk to node: diagonal/linear or both

	obstructions map[Node]struct{} //list of obstructions in the grid, not passable

	// todo (weights): potential problem of duplicating and lost elements for map[{0,0}]map[{0,1}] = 1; map[{0,1}]map[{0,0}] = 1,
	// the same movements and cost
	visited map[INode]struct{} //list of visited nodes
	start   Node               //start node
	target  Node               //target node

	draw IDraw
}

// InBound is it node located in the grid
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

// Neighbours get all neighbourhoods
func (g *SquareGrid) Neighbours(node INode) []Node {
	x, y := node.Position()
	nodes := make([]Node, 0, 8)

	if g.walkBehaviour == DiagonalWalk || g.walkBehaviour == AllWalk {
		nodes = append(nodes,
			Node{x-1, y-1},
			Node{x+1, y+1},
			Node{x+1, y-1},
			Node{x-1, y+1},
		)
	}

	if g.walkBehaviour == LinearWalk || g.walkBehaviour == AllWalk {
		nodes = append(nodes,
			Node{x, y-1},
			Node{x+1, y},
			Node{x, y+1},
			Node{x-1, y},
		)
	}

	return NodeFilter(nodes, func(n Node) bool {
		return g.Passable(n) && g.InBound(n)
	})
}

// Cost of movements from `node` to `node`
func (g *SquareGrid) Cost(from, to INode) uint32 {
	return 1
}

type IDraw interface {
	Point(x int, y int, pointType color.RGBA)
	Visit(x int, y int)
	Save()
}

// NewSquareGrid new instance of grid
func NewSquareGrid(width, height uint32, walkType WalkBehaviour, draw IDraw) *SquareGrid {
	log.Printf("Create rectangle grid, height: %v, width: %v", height, width)

	return &SquareGrid{
		width:         width,
		height:        height,
		obstructions:  map[Node]struct{}{},
		walkBehaviour: walkType,
		draw:          draw,
	}
}

// AddObstructions add list of walls as nodes
func (g *SquareGrid) AddObstructions(nodes ...Node) {
	for _, n := range nodes {
		if g.InBound(n) {
			g.obstructions[n] = struct{}{}
			continue
		}
	}
}

// AddWall add wall according on Heigt, Width
func (g *SquareGrid) AddWall(start Node, height, width int) {
	xPos, yPos := start.Position()
	for x := int(xPos); x < width+int(xPos); x++ {
		for y := int(yPos); y < height+int(yPos); y++ {
			g.AddObstructions(Node{uint32(x), uint32(y)})
			if g.draw != nil {
				g.draw.Point(x, y, PointWall)
			}
		}
	}
}

// String create preview of grid
func (g *SquareGrid) String() string {
	out := ""
	for y := 0; y < int(g.height); y++ {
		for x := 0; x < int(g.width); x++ {
			node := Node{uint32(x), uint32(y)}
			if g.target.Equal(node) {
				out += charTarget
				continue
			}
			if g.start.Equal(node) {
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

// SaveAnimation of current grid state
func (g *SquareGrid) SaveAnimation() {
	g.draw.Save()
}

// Visit set node of grid as visited
func (g *SquareGrid) Visit(node INode) {
	if g.visited == nil {
		g.visited = make(map[INode]struct{})
	}
	g.visited[node] = struct{}{}

	// for image draw
	if g.draw != nil {
		x, y := node.Position()
		g.draw.Visit(int(x), int(y))
	}
}

// Start set start node (only for draw)
func (g *SquareGrid) Start(node Node) {
	g.start = node

	//for image draw
	if g.draw != nil {
		x, y := node.Position()
		g.draw.Point(int(x), int(y), PointStart)
	}
}

// Target set start node (only for draw)
func (g *SquareGrid) Target(node Node) {
	g.target = node

	//for image draw
	if g.draw != nil {
		x, y := node.Position()
		g.draw.Point(int(x), int(y), PointFinish)
	}
}
