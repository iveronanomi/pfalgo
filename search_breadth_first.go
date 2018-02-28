package pfalgo

// BreadthFirstSearch ...
func BreadthFirstSearch(g *GridGraph, from, to INode,  callback func(graph *GridGraph, from, to INode)) (path map[INode]INode) {
	queue := NewQueue()
	path = map[INode]INode{from: nil}

	queue.Add(from)

	for queue.Len() > 0 {
		current := queue.Get()
		if current.Equal(to) {
			break
		}
		if callback != nil {
			callback(g, from, current)
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
