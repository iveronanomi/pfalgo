package main

import (
	"fmt"

	"github.com/iveronanomi/pfalgo/algo"
)

func main() {
	//Create Grid
	sg := algo.NewSquareGrid(30, 15, nil)
	//Add walls to grid
	sg.AddWall(algo.Node{3, 3}, 2, 9)
	sg.AddWall(algo.Node{13, 4}, 2, 11)
	sg.AddWall(algo.Node{21, 0}, 2, 7)
	sg.AddWall(algo.Node{21, 5}, 5, 2)

	fmt.Println(sg.String())

	cf := algo.BreadthFirstSearch2(sg, algo.Node{8, 7})

	fmt.Printf("\nCameFrom:\n%v\n", cf)
}
