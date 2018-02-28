package pfalgo

// INode interface
type INode interface {
	Position() (int, int) // get position of node
	Equal(node INode) bool // is equal nodes (check positions)
}

// IDraw interface
type IDrawer interface {
	Walk(x int, y int) //mark position as passed and add new frame
	Wall(x int, y int) //mark position as wall element
	Start(x int, y int) //mark position as start position
	Target(x int, y int) //mark position as target position
	Save() //save to file
}
