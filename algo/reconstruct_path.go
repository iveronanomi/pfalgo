package algo

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
