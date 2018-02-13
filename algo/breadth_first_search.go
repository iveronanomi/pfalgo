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
func BreadthFirstSearch2(g *SquareGrid, start Node) {
	f := NewQueue()
	f.Put(start)
	visited := map[uint32]map[uint32]bool{}
	visited[start[0]] = map[uint32]bool{
		start[1]: true,
	}

	for !f.Empty() {
		current := f.Get()
		log.Printf("Visited: %v", current)
		for _, next := range g.Neighbours(current) {
			if v, ok := visited[next[0]][next[1]]; !ok || !v {
				log.Printf("%v : %#v", current, f.nodes)
				f.Put(next)
				if _, ok := visited[next[0]]; !ok {
					visited[next[0]] = map[uint32]bool{
						next[1]: true,
					}
				} else {
					visited[next[0]][next[1]] = true
				}
			}
		}
	}
}
