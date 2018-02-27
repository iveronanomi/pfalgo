package pfalgo

// INode interface
type INode interface {
	Position() (int, int)
	Equal(node INode) bool
}

// IDraw interface
type IDraw interface {
	Walk(x int, y int)
	Wall(x int, y int)
	Start(x int, y int)
	Finish(x int, y int)
	Save(bool)
}
