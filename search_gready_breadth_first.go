package pfalgo

// GreedyBreadthFirstSearch ...
func GreedyBreadthFirstSearch(g *GridGraph, from, to INode, callback func(graph *GridGraph, from, to INode)) map[INode]INode {
	queue := NewPriorityQueue()
	queue.Add(from, 0)
	path := map[INode]INode{from: nil}

	for queue.Len() > 0 {
		current := queue.Get()
		if current.Equal(to) {
			break
		}
		if callback != nil {
			callback(g, from, to)
		}
		for _, next := range g.Neighbours(current) {
			if _, ok := path[next]; !ok {
				queue.Add(next, heuristic(to, next))
				path[next] = current
			}
		}
	}
	return path
}
