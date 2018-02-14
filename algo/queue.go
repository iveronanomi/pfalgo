package algo

// Queue stucture (actually it's a stack)
type Queue struct {
	nodes []Node
}

// IStack interface describes walking by nodes queue //todo
type IStack interface {
	IsEmpty() bool
	Put(n INode)
	Pop() INode
	Len() uint64
}

// NewQueue create new Queue
func NewQueue() *Queue {
	return &Queue{
		nodes: []Node{},
	}
}

// List get list of available nodes
func (q *Queue) List() []Node {
	return q.nodes
}

// PopLeft get first node an remove it from queue
func (q *Queue) PopLeft() Node {
	n := q.nodes[0]
	q.nodes = q.nodes[1:]
	return n
}

// Empty is it empty queue
func (q *Queue) Empty() bool {
	return len(q.nodes) < 1
}

// Put in queue
func (q *Queue) Put(n Node) {
	q.nodes = append(q.nodes, n)
}

// Get same as pop left
func (q *Queue) Get() Node {
	return q.PopLeft()
}
