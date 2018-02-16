package algo

import "log"

// Dijkstra search
func Dijkstra(g *SquareGrid, start, target INode) (map[INode]INode, map[INode]int) {
	q := PriorityQueue{}
	q.Push(start)
	cameFrom := map[INode]INode{
		start: nil,
	}
	costSoFar := map[INode]int{
		start: 0,
	}

	var itterations int //debug
	for q.Len() != 0 {
		itterations++ //debug
		item := q.Pop()
		current, ok := item.(INode)
		log.Printf("Current: %v", current)
		if !ok {
			log.Print("Could not conver queue item to INode")
		}
		// check is same position as target position
		if current.Equal(target) {
			break
		}
		g.Visit(current)
		for _, next := range g.Neighbours(current) {
			newCost := costSoFar[current] + g.Cost(current, next)
			// log.Printf("newCost: %d", newCost)
			if v, ok := costSoFar[next]; !ok || newCost < v {
				costSoFar[next] = newCost
				// priority := newCost
				q.Push(next) //set priorirt for node
				cameFrom[next] = current
			}
		}
	}
	log.Printf("Itterations %d", itterations) //debug
	return cameFrom, costSoFar
}
