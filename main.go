package main

import (
	"fmt"

	"github.com/iveronanomi/pfalgo/algo"
)

func main() {
	//Create Grid
	sg := algo.NewSquareGrid(30, 15, nil)

	//Add walls to grid
	sg.AddWall(algo.Node{3, 3}, 9, 2)
	sg.AddWall(algo.Node{13, 4}, 11, 2)
	sg.AddWall(algo.Node{21, 0}, 7, 2)
	sg.AddWall(algo.Node{21, 5}, 2, 5)

	cf := algo.BreadthFirstSearch3(sg, algo.Node{8, 7}, algo.Node{24, 0})

	fmt.Println(sg.String())
	fmt.Printf("\nCameFrom:\n%v\n", cf)
}
