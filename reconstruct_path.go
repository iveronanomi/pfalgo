package pfalgo

// ReconstructPath ...
func ReconstructPath(cameFrom map[INode]INode, from, to INode, reversed bool) (path []INode) {
	current := to
	path = []INode{}

	for !current.Equal(from) {
		path = append(path, current)
		current = cameFrom[current]
		if current == nil {
			break
		}
	}

	path = append(path, from)

	if reversed {
		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}
	}

	return path
}
