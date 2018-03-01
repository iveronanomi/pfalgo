package pfalgo

// DijkstraSearch ...
func DijkstraSearch(g *GridGraph, from, to INode, callback func(graph *GridGraph, from, to INode)) (path map[INode]INode, cost map[INode]int) {
	q := NewPriorityQueue()
	q.Add(from, 0)
	path = map[INode]INode{from: nil}
	cost = map[INode]int{from: 0}

	for q.Len() > 0 {
		current := q.Get()
		if current.Equal(to) {
			break
		}
		if callback != nil {
			callback(g, from, to)
		}
		for _, next := range g.Neighbours(current) {
			newCost := cost[current] + g.Cost(current, next)
			if v, ok := cost[next]; !ok || newCost < v {
				cost[next] = newCost
				q.Add(next, int(newCost))
				path[next] = current
			}
		}
	}
	return path, cost
}
