package algo

// AStar search
func AStar(g *SquareGrid, start, target INode) (map[INode]INode, map[INode]uint32) {
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
		// check is same position as target position
		if current.Equal(target) {
			break
		}
		//draw
		if !current.Equal(start) {
			g.Visit(current)
		}
		for _, next := range g.Neighbours(current) {
			newCost := costSoFar[current] + g.Cost(current, next)
			// log.Printf("newCost: %d", newCost)
			if v, ok := costSoFar[next]; !ok || newCost < v {
				costSoFar[next] = newCost
				q.Add(next, int(newCost)) //set priorirt for node
				cameFrom[next] = current
			}
		}
	}
	return cameFrom, costSoFar
}
