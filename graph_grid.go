package pfalgo

// WalkBehaviour ..
type WalkBehaviour uint8

const (
	// DiagonalWalk ...
	DiagonalWalk WalkBehaviour = 1 << iota
	// LinearWalk ...
	LinearWalk
	// AllWalk - walks all directions Diagonally and Linear
	AllWalk
)

// GridGraph space
type GridGraph struct {
	width         int           //width of grid
	height        int           //height of grid
	walkBehaviour WalkBehaviour //is it possible to walk to node: diagonal/linear or both

	obstructions map[int]map[int]struct{} //list of Obstructions in the grid, not passable
	visited      map[int]map[int]struct{} //list of Visited nodes
	start        [2]int                   //start node
	target       [2]int                   //target node

	Drawer IDrawer
}

// NewSquareGrid new instance of grid
func NewSquareGrid(width, height int, walkType WalkBehaviour, drawer IDrawer) *GridGraph {
	return &GridGraph{
		width:         width,
		height:        height,
		obstructions:  make(map[int]map[int]struct{}),
		visited:       make(map[int]map[int]struct{}),
		walkBehaviour: walkType,
		Drawer:        drawer,
	}
}

// InBound is it node located in the grid
func (g *GridGraph) InBound(x, y int) bool {
	return (0 <= x && x < int(g.width)) && (0 <= y && y < int(g.height))
}

// Passable is it passable node
func (g *GridGraph) Passable(x, y int) bool {
	_, ok := g.obstructions[x][y]
	return !ok
}

// NodeFilter function
func NodeFilter(nodes []Node, filter func(x, y int) bool) []Node {
	var result []Node
	for _, n := range nodes {
		if filter(n.X, n.Y) {
			result = append(result, n)
		}
	}
	return result
}

// Neighbours get all neighbourhoods
func (g *GridGraph) Neighbours(node INode) []Node {
	x, y := node.Position()
	nodes := make([]Node, 0, 8)

	if g.walkBehaviour == DiagonalWalk || g.walkBehaviour == AllWalk {
		nodes = append(nodes,
			Node{x - 1, y - 1},
			Node{x + 1, y + 1},
			Node{x + 1, y - 1},
			Node{x - 1, y + 1},
		)
	}

	if g.walkBehaviour == LinearWalk || g.walkBehaviour == AllWalk {
		nodes = append(nodes,
			Node{x, y - 1},
			Node{x + 1, y},
			Node{x, y + 1},
			Node{x - 1, y},
		)
	}

	return NodeFilter(nodes, func(x, y int) bool {
		return g.Passable(x, y) && g.InBound(x, y)
	})
}

// Cost of movements from `node` to `node`
func (g *GridGraph) Cost(from, to INode) int {
	return 1
}

// AddWall add wall according on Height, Width
func (g *GridGraph) AddWall(xPos, yPos, height, width int) {
	for x := int(xPos); x < width+int(xPos); x++ {
		for y := int(yPos); y < height+int(yPos); y++ {
			if !g.InBound(x, y) {
				continue
			}
			if _, ok := g.obstructions[x]; !ok {
				g.obstructions[x] = map[int]struct{}{}
			}
			g.obstructions[x][y] = struct{}{}
			if g.Drawer != nil {
				g.Drawer.Wall(x, y)
			}
		}
	}
}

// Visit set node of grid as Visited
func (g *GridGraph) Visit(x, y int) {
	if _, ok := g.visited[x]; !ok {
		g.visited[x] = map[int]struct{}{}
	}
	g.visited[x][y] = struct{}{}

	// for image Drawer
	if g.Drawer != nil {
		g.Drawer.Walk(x, y)
	}
}

// SetStart set start node (only for Drawer)
func (g *GridGraph) SetStart(x, y int) {
	g.start = [2]int{x, y}

	//for image Drawer
	if g.Drawer != nil {
		g.Drawer.Start(x, y)
	}
}

// SetTarget set start node (only for Drawer)
func (g *GridGraph) SetTarget(x, y int) {
	g.target = [2]int{x, y}

	//for image Drawer
	if g.Drawer != nil {
		g.Drawer.Target(x, y)
	}
}

// Start point
func (g *GridGraph) Start() (int, int) {
	return g.start[0], g.start[1]
}

// Target point
func (g *GridGraph) Target() (int, int) {
	return g.target[0], g.target[1]
}

// Visited point
func (g *GridGraph) Visited(x, y int) bool {
	_, ok := g.visited[x][y]
	return ok
}

//Height of graph
func (g *GridGraph) Height() int {
	return g.height
}

//Width of graph
func (g *GridGraph) Width() int {
	return g.width
}
