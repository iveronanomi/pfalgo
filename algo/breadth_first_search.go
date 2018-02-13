package algo

import (
	"log"
)

// BreadthFirstSearch ...
func BreadthFirstSearch(g SimpleGraph, start string) {
	f := NewQueue()
	f.Put(start)
	visited := map[string]bool{
		start: true,
	}

	for !f.Empty() {
		current := f.Get()
		log.Printf("Visited: %s", current)
		for _, next := range g.Neiborhood(current) {
			if v, ok := visited[next]; !ok || !v {
				f.Put(next)
				visited[next] = true
			}
		}
	}
}
