package graphWeight

// Graph struct
type Graph struct {
	vertices map[int]*Vertex
}
// 0 |  [int]*Edge --> int, Vertex
// 1 | 0 3
// 2 |
// 3 |

// Vertex struct
type Vertex struct {
	value int
	edges map[int]*Edge
}

// Edge struct
type Edge struct {
	weight int
	vertex *Vertex
}

// Add Vertex
func (g *Graph) AddVertex (key, value int){
	g.vertices[key] = &Vertex{value: value, edges: map[int]*Edge{}}
}
// Add Edge
func (g *Graph) AddEdge (w int, srck, desk int){
	if _, ok := g.vertices[srck]; !ok{
		return
	}
	if _, ok := g.vertices[desk]; !ok{
		return
	}
	g.vertices[srck].edges[desk] = &Edge{weight: w, vertex: g.vertices[desk]}
}
// NewGraph func
// Print func
