package algo

// Dijkstra search
func Dijkstra(g *SquareGrid, start, target INode) (map[INode]INode, map[INode]int) {
	return nil, nil

	// q := PriorityQueue{}
	// q.Push(start)
	// cameFrom := map[INode]INode{
	// 	start: nil,
	// }
	// costSoFar := map[INode]int{
	// 	start: 0,
	// }

	// for q.Len() != 0 {
	// 	item := q.Pop()
	// 	current, ok := item.(INode)
	// 	if !ok {
	// 		log.Print("Could not conver queue item to INode")
	// 	}
	// 	// check is same position as target position
	// 	if current.Equal(target) {
	// 		break
	// 	}
	// 	for _, next := range g.Neighbours(current) {
	// 		newCost := costSoFar[current] + g.Cost(current, next)
	// 		if v, ok := costSoFar[next]; !ok || newCost < v {
	// 			costSoFar[next] = newCost
	// 			// priority := newCost
	// 			q.Push(next) //set priorirt for node
	// 			cameFrom[next] = current
	// 		}
	// 	}
	// }
	// return cameFrom, costSoFar
}
