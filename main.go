package main

import (
	"fmt"

	"github.com/iveronanomi/pfalgo/algo"
)

func main() {

	sg, start, end := getGrid()
	// algo.BreadthFirstSearch(sg, start, end)

	algo.GreadyBreadthFirstSearch(sg, start, end)

	// algo.Dijkstra(sg, start, end)

	// rp := algo.ReconstructPath(cp, start, end, true)
	// fmt.Println(cp)
	// fmt.Println(rp)

	fmt.Println(sg.String())
	sg.SaveAnimation()
}

func getGrid() (*algo.SquareGrid, algo.Node, algo.Node) {
	//Create Grid
	sg := algo.NewSquareGrid(30, 15, algo.LinearWalk)

	start := algo.NewNode(24, 12, 1)
	end := algo.NewNode(24, 1, 1)
	//Add walls to grid
	sg.AddWall(algo.NewNode(3, 3, 1), 9, 2)
	sg.AddWall(algo.NewNode(13, 4, 1), 11, 2)
	sg.AddWall(algo.NewNode(21, 0, 1), 7, 2)
	sg.AddWall(algo.NewNode(21, 5, 1), 2, 5)

	sg.Start(start) //only for a draw
	sg.Target(end)  //only for a draw
	return sg, start, end
}
