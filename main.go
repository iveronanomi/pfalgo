package main

import (
	"github.com/iveronanomi/pfalgo/algo"
)

func main() {
	//Create Grid
	sg := algo.NewSquareGrid(30, 15, algo.AllWalk)

	start := algo.NewNode(24, 12, 0)
	end := algo.NewNode(24, 1, 0)
	//Add walls to grid
	sg.AddWall(algo.NewNode(3, 3, 0), 9, 2)
	sg.AddWall(algo.NewNode(13, 4, 0), 11, 2)
	sg.AddWall(algo.NewNode(21, 0, 0), 7, 2)
	sg.AddWall(algo.NewNode(21, 5, 0), 2, 5)

	sg.Start(start) //only for a draw
	sg.Target(end)  //only for a draw

	// algo.BreadthFirstSearch(sg, algo.NewNode(24, 12, 0), algo.NewNode(24, 0, 0))
	// fmt.Println(sg.String())

	algo.GreadyBreadthFirstSearch(sg, start, end)
	// fmt.Println(sg.String())
	sg.SaveAnimation()

	// algo.Dijkstra(sg, start, end)
	// fmt.Println(sg.String())
	// rp := algo.ReconstructPath(cp, start, end, true)
	// fmt.Println(cp)
	// fmt.Println(rp)
}
