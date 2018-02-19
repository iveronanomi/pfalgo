package algo

import (
	"container/heap"
	"log"
)

// Dijkstra search
func Dijkstra(g *SquareGrid, start, target INode) (map[INode]INode, map[INode]int) {
	q := &PriorityQueue{}
	heap.Init(q)
	q.Push(start)
	cameFrom := map[INode]INode{
		start: nil,
	}
	costSoFar := map[INode]int{
		start: 0,
	}

	var itterations int //debug
	for q.Len() > 0 {
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
				q.PriorityPush(next, newCost) //set priorirt for node
				cameFrom[next] = current
			}
		}
	}
	log.Printf("Itterations %d", itterations) //debug
	return cameFrom, costSoFar
}

// ReversePath ...
func ReversePath(cameFrom map[INode]INode, start, target INode) []INode {
	current := target
	path := []INode{}
	for !current.Equal(start) {
		path = append(path, current)
		current = cameFrom[current]
	}
	path = append(path, start)
	return path
}
