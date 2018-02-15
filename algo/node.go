package algo

type nodeStatus uint16

const (
	statusVisited nodeStatus = 1 << iota
	statusNotPassable
	statusTarget
	statusStart
)

// INode intarface
type INode interface {
	Position() (uint32, uint32)
	Equal(node INode) bool
}

// Node structure
type Node struct {
	pos    [2]uint32
	weight int32
	status nodeStatus
}

// NewNode create instance of node
func NewNode(x, y uint32, weight int32) *Node {
	return &Node{
		pos:    [2]uint32{x, y},
		weight: weight,
	}
}

// Position get position of node (x, y)
func (n *Node) Position() (uint32, uint32) {
	return n.pos[0], n.pos[1]
}

// Equal with given INode if positions are equals
func (n *Node) Equal(node INode) bool {
	x1, y1 := n.Position()
	x2, y2 := node.Position()
	return x1 == x2 && y1 == y2
}

// setFlag ...
func (n *Node) setFlag(f nodeStatus) {
	n.status |= f
}

// hasFlag ...
func (n *Node) hasFlag(f nodeStatus) bool {
	return (f)&n.status == f
}

// removeFlag ...
func (n *Node) removeFlag(f nodeStatus) {
	n.status ^= f
}

// Visit set node as visited node
func (n *Node) Visit() *Node {
	n.setFlag(statusVisited)
	return n
}

// Target set node as target node
func (n *Node) Target() *Node {
	n.setFlag(statusTarget)
	return n
}

// Start set node as start node
func (n *Node) Start() *Node {
	n.setFlag(statusStart)
	return n
}

// NotPassable set node as not passable node
func (n *Node) NotPassable() *Node {
	n.setFlag(statusNotPassable)
	return n
}

// IsVisited ...
func (n *Node) IsVisited() bool {
	return n.hasFlag(statusVisited)
}

// IsPassable ...
func (n *Node) IsPassable() bool {
	return !n.hasFlag(statusNotPassable)
}

// IsTarget ...
func (n *Node) IsTarget() bool {
	return n.hasFlag(statusTarget)
}

// IsStart ...
func (n *Node) IsStart() bool {
	return n.hasFlag(statusStart)
}
