package main

import (
	"fmt"

	"github.com/iveronanomi/pfalgo/algo"
)

func main() {
	sg, start, end := getGrid()
	//cf := algo.BreadthFirstSearch(sg, start, end)
	cf := algo.GreedyBreadthFirstSearch(sg, start, end)
	//cf, _ := algo.DijkstraSearch(sg, start, end)
	//cf, _ := algo.AStarSearch(sg, start, end)
	fmt.Println(algo.String(sg))
	rp := algo.ReconstructPath(cf, start, end, false)
	drawReconstructedPath(rp)
}

func getGrid() (*algo.SquareGridGraph, algo.Node, algo.Node) {
	//Create Grid
	w, h := 30, 15
	sg := algo.NewSquareGrid(
		uint32(w),
		uint32(h),
		algo.LinearWalk,
		algo.NewGifGraph(w, h, "out/out.gif"),
	)
	start := algo.Node{24, 12}
	end := algo.Node{0, 10}

	//Add walls to grid
	sg.AddWall(algo.Node{3, 3}, 9, 2)
	sg.AddWall(algo.Node{13, 4}, 11, 2)
	sg.AddWall(algo.Node{21, 0}, 7, 2)
	sg.AddWall(algo.Node{21, 5}, 2, 5)
	sg.Start(start) //only for a draw
	sg.Target(end)  //only for a draw

	return sg, start, end
}

func drawReconstructedPath(rp []algo.INode) {
	sg, start, end := getGrid()
	for _, v := range rp {
		if v.Equal(start) || v.Equal(end) {
			continue
		}
		sg.Visit(v)
	}
	fmt.Println(algo.String(sg))
}
