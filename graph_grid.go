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
	width         uint32        //width of grid
	height        uint32        //height of grid
	walkBehaviour WalkBehaviour //is it possible to walk to node: diagonal/linear or both

	obstructions map[int]map[int]struct{}  //list of obstructions in the grid, not passable
	visited      map[int]map[int]struct{} //list of visited nodes
	start        [2]int               //start node
	target       [2]int               //target node

	Drawer IDrawer
}

// NewSquareGrid new instance of grid
func NewSquareGrid(width, height uint32, walkType WalkBehaviour, drawer IDrawer) *GridGraph {
	return &GridGraph{
		width:         width,
		height:        height,
		obstructions:  map[int]map[int]struct{}{},
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
func (g *GridGraph) Cost(from, to INode) uint32 {
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

// Visit set node of grid as visited
func (g *GridGraph) Visit(x, y int) {
	if g.visited == nil {
		g.visited = make(map[int]map[int]struct{})
	}
	if _, ok := g.visited[x]; !ok {
		g.visited[x] = map[int]struct{}{}
	}
	g.visited[x][y] = struct{}{}

	// for image Drawer
	if g.Drawer != nil {
		g.Drawer.Walk(x, y)
	}
}

// Start set start node (only for Drawer)
func (g *GridGraph) Start(x, y int) {
	g.start = [2]int{x,y}

	//for image Drawer
	if g.Drawer != nil {
		g.Drawer.Start(x, y)
	}
}

// Target set start node (only for Drawer)
func (g *GridGraph) Target(x, y int) {
	g.target = [2]int{x,y}

	//for image Drawer
	if g.Drawer != nil {
		g.Drawer.Finish(x, y)
	}
}
