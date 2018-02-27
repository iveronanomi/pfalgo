package pfalgo

import "math"

func heuristic(a, b INode) int {
	ax, ay := a.Position()
	bx, by := b.Position()
	result := math.Abs(float64(ax)-float64(bx)) + math.Abs(float64(ay)-float64(by))
	return int(result)
}
