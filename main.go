package main

import (
	"fmt"

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

	// algo.BreadthFirst(sg, algo.NewNode(24, 12, 0), algo.NewNode(24, 0, 0))

	// fmt.Println(sg.String())

	algo.Dijkstra(sg, start, end)

	fmt.Println(sg.String())
}
