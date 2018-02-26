package algo


const (
	charObstruction = " ◼"
	charGrid        = " ◻"
	charStart       = " Ⓐ"
	charTarget      = " Ⓩ"
	charVisited     = " ◆"
)

// String create preview of grid
func String(g *SquareGridGraph) string {
	out := ""
	for y := 0; y < int(g.height); y++ {
		for x := 0; x < int(g.width); x++ {
			node := Node{uint32(x), uint32(y)}
			if g.target.Equal(node) {
				out += charTarget
				continue
			}
			if g.start.Equal(node) {
				out += charStart
				continue
			}
			if !g.Passable(node) {
				out += charObstruction
				continue
			}
			if _, ok := g.visited[node]; ok {
				out += charVisited
				continue
			}
			out += charGrid
		}
		out += "\n\r"
	}
	return out
}
