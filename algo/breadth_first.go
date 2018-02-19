package algo

// BreadthFirstSearch ...
func BreadthFirstSearch(g *SquareGrid, start, target INode) map[INode]INode {
	queue := NewQueue()

	queue.Put(start)
	cameFrom := map[INode]INode{
		start: nil,
	}
	for !queue.Empty() {
		current := queue.Get()
		if current.Equal(target) {
			break
		}

		g.Visit(current) //debug: mark graph node as visited

		for _, next := range g.Neighbours(current) {
			if _, ok := cameFrom[next]; !ok {
				cameFrom[next] = current
				queue.Put(next)
			}
		}
	}
	return cameFrom
}
