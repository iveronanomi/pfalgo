package pfalgo

// AStarSearch ...
func AStarSearch(g *GridGraph, from, to INode, callback func(graph *GridGraph, from, to INode)) (path map[INode]INode, cost map[INode]uint32) {
	queue := NewPriorityQueue()
	queue.Add(from, 0)
	path = map[INode]INode{from: nil}
	cost = map[INode]uint32{from: 0}

	for queue.Len() > 0 {
		current := queue.Get()
		if current.Equal(to) {
			break
		}
		if callback != nil {
			callback(g, from, current)
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
