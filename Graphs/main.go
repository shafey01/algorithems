package main

import (
	"github.com/shafey01/algorithems/Graphs/graph"
)

func main() {
	graph := graph.NewGraph()

	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(6, 2)
	graph.AddEdge(3, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(1, 5)

	graph.Print()
}
