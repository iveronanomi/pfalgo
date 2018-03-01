package pfalgo

// GridGraphSymbols grid graph symbol type
type GridGraphSymbols int8

const (
	//SymbolObstruction obstruction node symbol
	SymbolObstruction GridGraphSymbols = iota << 1
	//SymbolGrid empty grid node symbol
	SymbolGrid
	//SymbolStart start node symbol
	SymbolStart
	//SymbolTarget target node symbol
	SymbolTarget
	//SymbolVisited symbol of visited node
	SymbolVisited
)

var gridStringerSymbols = map[GridGraphSymbols]string{
	SymbolObstruction: " ◼",
	SymbolGrid:        " ◻",
	SymbolStart:       " Ⓐ",
	SymbolTarget:      " Ⓩ",
	SymbolVisited:     " ◆",
}

// String create string grid visualisation
func String(g IGridGraph, symbols map[GridGraphSymbols]string) string {
	if symbols == nil {
		symbols = gridStringerSymbols
	}
	var out string
	startX, startY := g.Start()
	targetX, targetY := g.Target()

	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if targetX == x && targetY == y {
				out += symbols[SymbolTarget]
				continue
			}
			if startX == x && startY == y {
				out += symbols[SymbolStart]
				continue
			}
			if !g.Passable(x, y) {
				out += symbols[SymbolObstruction]
				continue
			}
			if g.Visited(x, y) {
				out += symbols[SymbolVisited]
				continue
			}
			out += symbols[SymbolGrid]
		}
		out += "\n\r"
	}
	return out
}
