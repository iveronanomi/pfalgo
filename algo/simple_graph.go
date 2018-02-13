package algo

// SimpleGraph ...
type SimpleGraph struct {
	Nodes map[string][]string
}

// Neiborhood ...
func (g *SimpleGraph) Neiborhood(n string) []string {
	return g.Nodes[n]
}
