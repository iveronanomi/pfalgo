package algo

// Graph ...
type Graph struct {
	Nodes map[string][]string
}

// Neighborhood ...
func (g *Graph) Neighborhood(n string) []string {
	return g.Nodes[n]
}
