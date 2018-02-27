package pfalgo

// Node structure
type Node struct {
	X int
	Y int
}

// Position get position of node (x, y)
func (n Node) Position() (int, int) {
	return n.X, n.Y
}

// Equal with given INode if positions are equals
func (n Node) Equal(node INode) bool {
	x, y := node.Position()
	return n.X == x && n.Y == y
}
