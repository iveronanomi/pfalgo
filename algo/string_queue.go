package algo

// StringQueue structure
type StringQueue struct {
	nodes []string
}

// NewStringQueue create new StringQueue
func NewStringQueue() *StringQueue {
	return &StringQueue{
		nodes: []string{},
	}
}

// PopLeft get first node an remove it from queue
func (q *StringQueue) PopLeft() string {
	n := q.nodes[0]
	q.nodes = q.nodes[1:]
	return n
}

// Empty is it empty queue
func (q *StringQueue) Empty() bool {
	return len(q.nodes) < 1
}

// Put in queue
func (q *StringQueue) Put(n string) {
	q.nodes = append(q.nodes, n)
}

// Get same as pop left
func (q *StringQueue) Get() string {
	return q.PopLeft()
}
