package algo

// GreedyBreadthFirstSearch ...
func GreedyBreadthFirstSearch(g *SquareGrid, from, to INode) map[INode]INode {
	queue := NewPriorityQueue()
	queue.Add(from, 0)
	path := map[INode]INode{from: nil}

	for queue.Len() > 0 {
		current := queue.Get()

		if current.Equal(to) {
			break
		}

		if !current.Equal(from) {
			g.Visit(current) //debug: mark graph node as visited
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
