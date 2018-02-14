package algo

// Node of square grid
type Node [2]uint32

// NodeFlag ...
type NodeFlag uint16

const (
	flagVisited NodeFlag = 1 << iota
	flagNotPassable
	flagTarget
	flagStart
)

// INode intarface
type INode interface {
	Position() (uint32, uint32)
}

// Position get position of node (x, y)
func (n *Node) Position() (uint32, uint32) {
	return n[0], n[1]
}

// Position get position of node (x, y)
func (n *AdvancedNode) Position() (uint32, uint32) {
	return n.position[0], n.position[1]
}

// AdvancedNode todo
type AdvancedNode struct {
	position [2]uint32
	weight   float32
	flags    NodeFlag
}

// setFlag ...
func (n *AdvancedNode) setFlag(f NodeFlag) {
	n.flags |= f
}

// hasFlag ...
func (n *AdvancedNode) hasFlag(f NodeFlag) bool {
	return (f)&n.flags == f
}

// removeFlag ...
func (n *AdvancedNode) removeFlag(f NodeFlag) {
	n.flags ^= f
}

// Visit set node as visited node
func (n *AdvancedNode) Visit() *AdvancedNode {
	n.setFlag(flagVisited)
	return n
}

// Target set node as target node
func (n *AdvancedNode) Target() *AdvancedNode {
	n.setFlag(flagTarget)
	return n
}

// Start set node as start node
func (n *AdvancedNode) Start() *AdvancedNode {
	n.setFlag(flagStart)
	return n
}

// NotPassable set node as not passable node
func (n *AdvancedNode) NotPassable() *AdvancedNode {
	n.setFlag(flagNotPassable)
	return n
}

// IsVisited ...
func (n *AdvancedNode) IsVisited() bool {
	return n.hasFlag(flagVisited)
}

// IsPassable ...
func (n *AdvancedNode) IsPassable() bool {
	return !n.hasFlag(flagNotPassable)
}

// IsTarget ...
func (n *AdvancedNode) IsTarget() bool {
	return n.hasFlag(flagTarget)
}

// IsStart ...
func (n *AdvancedNode) IsStart() bool {
	return n.hasFlag(flagStart)
}
