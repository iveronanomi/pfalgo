package algo

// AStarSearch ...
func AStarSearch(g *SquareGridGraph, start, target INode) (path map[INode]INode, cost map[INode]uint32) {
	queue := NewPriorityQueue()
	queue.Add(start, 0)
	path = map[INode]INode{start: nil}
	cost = map[INode]uint32{start: 0}

	for queue.Len() > 0 {
		current := queue.Get()
		if current.Equal(target) {
			break
		}

		//mark graph node as visited
		if !current.Equal(start) {
			g.Visit(current)
		}

		for _, next := range g.Neighbours(current) {
			newCost := cost[current] + g.Cost(current, next)
			if v, ok := cost[next]; !ok || newCost < v {
				cost[next] = newCost
				queue.Add(next, int(newCost))
				path[next] = current
			}
		}
	}
	return path, cost
}
