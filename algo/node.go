package algo

// INode intarface
type INode interface {
	Position() (uint32, uint32)
	Equal(node INode) bool
	Cost() uint32
}

// Node structure
type Node struct {
	pos      [2]uint32
	moveCost uint32
}

// NewNode create instance of node
func NewNode(x, y, moveCost uint32) Node {
	return Node{
		pos:      [2]uint32{x, y},
		moveCost: moveCost,
	}
}

// Position get position of node (x, y)
func (n Node) Position() (uint32, uint32) {
	return n.pos[0], n.pos[1]
}

// Equal with given INode if positions are equals
func (n Node) Equal(node INode) bool {
	x1, y1 := n.Position()
	x2, y2 := node.Position()
	return x1 == x2 && y1 == y2
}

// Cost of movement
func (n Node) Cost() uint32 {
	return n.moveCost
}
