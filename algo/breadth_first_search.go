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
		// log.Printf("Visited: %s", current)
		for _, next := range g.Neiborhood(current) {
			if v, ok := visited[next]; !ok || !v {
				// log.Printf("%v : %#v", current, f.nodes)
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
		g.Visit(current) //draw
		// log.Printf("Came From: %v", current)
		for _, next := range g.Neighbours(current) {
			if v, ok := cameFrom[next]; !ok || !v {
				// log.Printf("%v : %#v", current, frontier.nodes)
				frontier.Put(next)
				cameFrom[next] = true
			}
		}
	}
	return cameFrom
}

// BreadthFirstSearch3 ...
func BreadthFirstSearch3(g *SquareGrid, start, target Node) {
	queue := NewQueue()
	queue.Put(start)
	cameFrom := map[Node]struct{}{}
	cameFrom[start] = struct{}{}

	log.Printf("Start node: %v", start)   //debug
	log.Printf("Target node: %v", target) //debug

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
	log.Printf("Came from : %v", cameFrom)
}

// DijkstraSearch ...
func DijkstraSearch(g *SquareGrid, start, target Node) {

}
