package pfalgo

const (
	charObstruction = " ◼"
	charGrid        = " ◻"
	charStart       = " Ⓐ"
	charTarget      = " Ⓩ"
	charVisited     = " ◆"
)

// String create string grid visualisation
func String(g *GridGraph) string {
	out := ""
	for y := 0; y < int(g.height); y++ {
		for x := 0; x < int(g.width); x++ {
			if g.target[0] == x && g.target[1] == y {
				out += charTarget
				continue
			}
			if g.start[0] == x && g.start[1] == y {
				out += charStart
				continue
			}
			if !g.Passable(x, y) {
				out += charObstruction
				continue
			}
			if _, ok := g.Visited[x][y]; ok {
				out += charVisited
				continue
			}
			out += charGrid
		}
		out += "\n\r"
	}
	return out
}
