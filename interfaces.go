package pfalgo

// INode interface
type INode interface {
	Position() (int, int)  // get position of node
	Equal(node INode) bool // is equal nodes (check positions)
}

// IDrawer interface
type IDrawer interface {
	Walk(x int, y int)   //mark position as passed and add new frame
	Wall(x int, y int)   //mark position as wall element
	Start(x int, y int)  //mark position as start position
	Target(x int, y int) //mark position as target position
	Save()               //save to file
}

// IGridGraph grid graph interface
type IGridGraph interface {
	InBound(x, y int) bool        // is it point belong to grid ot not
	Passable(x, y int) bool       // is it possible to walk throw this point
	Neighbours(node INode) []Node // neightbours of given node, depends on WalkBehaviour
	Cost(from, to INode) int      // cost of movements from node to nde
	Width() int                   //width of grid
	Height() int                  //height of grid
	Target() (x, y int)           // give a position of target node
	Start() (x, y int)            // give a position of start node
	Visited(x, y int) bool        //is it point visited or not
}
