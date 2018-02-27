package pfalgo

// Queue ...
type Queue struct {
	nodes []INode
}

// NewQueue create new Queue
func NewQueue() *Queue {
	return &Queue{}
}

// Add in queue
func (q *Queue) Add(n INode) {
	q.nodes = append(q.nodes, n)
}

// Get same as pop left
func (q *Queue) Get() INode {
	n := q.nodes[0]
	q.nodes = q.nodes[1:]
	return n
}

// Len of queue
func (q *Queue) Len() int {
	return len(q.nodes)
}
