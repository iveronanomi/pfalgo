package algo

// BreadthFirst ...
func BreadthFirst(g *SquareGrid, start, target Node) {
	queue := NewQueue()
	queue.Put(start)
	cameFrom := map[Node]struct{}{}
	cameFrom[start] = struct{}{}

	// log.Printf("Start node: %v", start)   //debug
	// log.Printf("Target node: %v", target) //debug

	for !queue.Empty() {
		// log.Printf("QUEUE: %v", queue.List()) //debug
		current := queue.Get()
		// log.Printf("Current node : %v", current) //debug

		if current == target {
			// log.Print("Goal has reached") //debug
			break
		}

		g.Visit(current) //draw and debug

		for _, next := range g.Neighbours(current) {
			if _, ok := cameFrom[next]; !ok {
				cameFrom[next] = struct{}{}
				queue.Put(next)
			}
		}
		// log.Printf("Came from : %v", cameFrom)
	}
	// log.Printf("Came from : %v", cameFrom)
}
