package pfalgo

import "container/heap"

// An Item is something we manage in a priority queue.
type Item struct {
	value    INode // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds NodeItems.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// Swap items in queue
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push item to queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop item from queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0: n-1]
	return item
}

// NewPriorityQueue create new priority queue
func NewPriorityQueue() *PriorityQueue {
	q := &PriorityQueue{}
	heap.Init(q)
	return q
}

// Add node item to queue
func (pq *PriorityQueue) Add(node INode, priority int) {
	item := &Item{
		value:    node,
		priority: priority,
	}
	heap.Push(pq, item)
}

// Get node element
func (pq *PriorityQueue) Get() INode {
	item := heap.Pop(pq).(*Item)
	return item.value
}
