package algo

// Dijkstra search
func Dijkstra(g *SquareGrid, start, target INode) (map[INode]INode, map[INode]uint32) {
	q := NewPriorityQueue()
	q.Add(start, 0)
	cameFrom := map[INode]INode{
		start: nil,
	}
	costSoFar := map[INode]uint32{
		start: 0,
	}

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
			newCost := costSoFar[current] + g.Cost(current, next)
			// log.Printf("newCost: %d", newCost)
			if v, ok := costSoFar[next]; !ok || newCost < v {
				costSoFar[next] = newCost
				// priority := newCost
				q.Add(next, int(newCost)) //set priorirt for node
				cameFrom[next] = current
			}
		}
	}
	return cameFrom, costSoFar
}

// ReconstructPath ...
func ReconstructPath(cameFrom map[INode]INode, start, target INode, reversed bool) []INode {
	current := target
	path := []INode{}
	for !current.Equal(start) {
		path = append(path, current)
		current = cameFrom[current]
	}
	path = append(path, start)
	if !reversed {
		return path
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
