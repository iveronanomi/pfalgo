package algo

import (
	"container/heap"
	"log"
	"math"
)

// GreadyBreadthFirstSearch ...
func GreadyBreadthFirstSearch(g *SquareGrid, start, target INode) map[INode]INode {
	queue := &PriorityQueue{}
	heap.Init(queue)
	queue.Push(start)
	cameFrom := map[INode]INode{
		start: nil,
	}

	for queue.Len() > 0 {
		e := queue.Pop()
		current := e.(INode)
		if current.Equal(target) {
			break
		}

		g.Visit(current) //debug: mark graph node as visited
		neighbours := g.Neighbours(current)
		for _, next := range neighbours {
			if _, ok := cameFrom[next]; !ok {
				priority := heuristic(target, next)
				log.Printf("next: %v, priorirt: %v", next, priority)
				cameFrom[next] = current
				queue.PriorityPush(next, priority)
			}
		}
	}
	return cameFrom
}

func heuristic(a, b INode) int {
	ax, ay := a.Position()
	bx, by := b.Position()
	result := math.Abs(float64(ax)-float64(bx)) + math.Abs(float64(ay)-float64(by))
	return int(result)
}
