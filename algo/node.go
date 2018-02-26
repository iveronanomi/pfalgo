package algo

// Node structure
type Node struct {
	X     uint32
	Y     uint32
}

// Position get position of node (x, y)
func (n Node) Position() (uint32, uint32) {
	return n.X, n.Y
}

// Equal with given INode if positions are equals
func (n Node) Equal(node INode) bool {
	x, y := node.Position()
	return n.X == x && n.Y == y
}

