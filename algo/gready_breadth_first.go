package algo

import (
	"math"
)

// GreadyBreadthFirstSearch ...
func GreadyBreadthFirstSearch(g *SquareGrid, start, target INode) map[INode]INode {
	queue := NewPriorityQueue()
	queue.Add(start, 0)
	cameFrom := map[INode]INode{start: nil}

	for queue.Len() > 0 {
		current := queue.Get()
		if current.Equal(target) {
			break
		}

		if !current.Equal(start) {
			g.Visit(current) //debug: mark graph node as visited
		}

		for _, next := range g.Neighbours(current) {
			if _, ok := cameFrom[next]; !ok {
				// log.Printf("next: %v, priority: %v", next, priority)
				queue.Add(next, heuristic(target, next))
				cameFrom[next] = current
			}
		}
	}
	return cameFrom
}

func heuristic(a, b INode) int {
	ax, ay := a.Position()
	bx, by := b.Position()
	result := math.Abs(
		float64(ax)-float64(bx)) + math.Abs(float64(ay)-float64(by))
	return int(result)
}
