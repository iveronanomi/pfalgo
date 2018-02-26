package algo

// Dijkstra search
func Dijkstra(g *SquareGrid, start, target INode) (path map[INode]INode, cost map[INode]uint32) {
	q := NewPriorityQueue()
	q.Add(start, 0)
	path = map[INode]INode{start: nil}
	cost = map[INode]uint32{start: 0}

	for q.Len() > 0 {
		current := q.Get()
		if current.Equal(target) {
			break
		}

		// draw option
		if !current.Equal(start) {
			g.Visit(current)
		}

		for _, next := range g.Neighbours(current) {
			newCost := cost[current] + g.Cost(current, next)
			// log.Printf("newCost: %d", newCost)
			if v, ok := cost[next]; !ok || newCost < v {
				cost[next] = newCost
				// priority := newCost
				q.Add(next, int(newCost)) //set priorirt for node
				path[next] = current
			}
		}
	}
	return path, cost
}
