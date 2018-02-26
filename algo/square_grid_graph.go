package algo

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

// SquareGridGraph space
type SquareGridGraph struct {
	width         uint32        //width of grid
	height        uint32        //height of grid
	walkBehaviour WalkBehaviour //is it possible to walk to node: diagonal/linear or both

	obstructions map[Node]struct{} //list of obstructions in the grid, not passable
	visited map[INode]struct{} //list of visited nodes
	start   Node               //start node
	target  Node               //target node

	draw IDraw
}

// NewSquareGrid new instance of grid
func NewSquareGrid(width, height uint32, walkType WalkBehaviour, draw IDraw) *SquareGridGraph {
	return &SquareGridGraph{
		width:         width,
		height:        height,
		obstructions:  map[Node]struct{}{},
		walkBehaviour: walkType,
		draw:          draw,
	}
}

// InBound is it node located in the grid
func (g *SquareGridGraph) InBound(node Node) bool {
	x, y := node.Position()
	return (0 <= x && x < g.width) && (0 <= y && y < g.height)
}

// Passable is it passable node
func (g *SquareGridGraph) Passable(node Node) bool {
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
func (g *SquareGridGraph) Neighbours(node INode) []Node {
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
func (g *SquareGridGraph) Cost(from, to INode) uint32 {
	return 1
}

// AddObstructions add list of walls as nodes
func (g *SquareGridGraph) AddObstructions(nodes ...Node) {
	for _, n := range nodes {
		if g.InBound(n) {
			g.obstructions[n] = struct{}{}
			continue
		}
	}
}

// AddWall add wall according on Heigt, Width
func (g *SquareGridGraph) AddWall(start Node, height, width int) {
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

// Visit set node of grid as visited
func (g *SquareGridGraph) Visit(node INode) {
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
func (g *SquareGridGraph) Start(node Node) {
	g.start = node

	//for image draw
	if g.draw != nil {
		x, y := node.Position()
		g.draw.Point(int(x), int(y), PointStart)
	}
}

// Target set start node (only for draw)
func (g *SquareGridGraph) Target(node Node) {
	g.target = node

	//for image draw
	if g.draw != nil {
		x, y := node.Position()
		g.draw.Point(int(x), int(y), PointFinish)
	}
}
