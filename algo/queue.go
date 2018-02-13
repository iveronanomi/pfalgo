package algo

// Queue stucture
type Queue struct {
	nodes []string
}

// NewQueue create new Queue
func NewQueue() *Queue {
	return &Queue{
		nodes: []string{},
	}
}

// PopLeft get first node an remove it from queue
func (q *Queue) PopLeft() string {
	n := q.nodes[0]
	q.nodes = q.nodes[1:]
	return n
}

// Empty is it empty queue
func (q *Queue) Empty() bool {
	return len(q.nodes) < 1
}

// Put in queue
func (q *Queue) Put(n string) {
	q.nodes = append(q.nodes, n)
}

// Get same as pop left
func (q *Queue) Get() string {
	return q.PopLeft()
}
