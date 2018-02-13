package algo

import (
	"log"
)

// BreadthFirstSearch1 ...
func BreadthFirstSearch1(g SimpleGraph, start string) {
	f := NewStringQueue()
	f.Put(start)
	visited := map[string]bool{
		start: true,
	}

	for !f.Empty() {
		current := f.Get()
		log.Printf("Visited: %s", current)
		for _, next := range g.Neiborhood(current) {
			if v, ok := visited[next]; !ok || !v {
				log.Printf("%v : %#v", current, f.nodes)
				f.Put(next)
				visited[next] = true
			}
		}
	}
}

// BreadthFirstSearch2 ...
func BreadthFirstSearch2(g *SquareGrid, start Node) map[Node]bool {
	frontier := NewQueue()
	frontier.Put(start)
	cameFrom := map[Node]bool{}
	cameFrom[start] = false

	for !frontier.Empty() {
		current := frontier.Get()
		log.Printf("Came From: %v", current)
		for _, next := range g.Neighbours(current) {
			if v, ok := cameFrom[next]; !ok || !v {
				log.Printf("%v : %#v", current, frontier.nodes)
				frontier.Put(next)
				cameFrom[next] = true
			}
		}
	}
	return cameFrom
}
