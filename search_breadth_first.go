package pfalgo

// BreadthFirstSearch ...
func BreadthFirstSearch(g *GridGraph, from, to INode) (path map[INode]INode) {
	queue := NewQueue()
	path = map[INode]INode{from: nil}

	queue.Add(from)

	for queue.Len() > 0 {
		current := queue.Get()
		if current.Equal(to) {
			break
		}

		//mark graph node as visited
		if !current.Equal(from) {
			g.Visit(current.Position())
		}

		for _, next := range g.Neighbours(current) {
			if _, ok := path[next]; !ok {
				path[next] = current
				queue.Add(next)
			}
		}
	}
	return path
}