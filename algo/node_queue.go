package algo

// Queue stucture
type Queue struct {
	nodes []Node
}

// NewQueue create new Queue
func NewQueue() *Queue {
	return &Queue{
		nodes: []Node{},
	}
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
