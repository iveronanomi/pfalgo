package main

import (
	"fmt"

	"github.com/iveronanomi/pfalgo/algo"
)

func main() {
	walls := map[uint32]map[uint32]struct{}{
		2: {
			2: struct{}{},
			3: struct{}{},
			4: struct{}{},
			5: struct{}{},
			6: struct{}{},
			7: struct{}{},
		},
	}

	sg := algo.NewSquareGrid(15, 5, walls)
	fmt.Println(sg.String())

	algo.BreadthFirstSearch2(sg, algo.Node{3, 5})
}
