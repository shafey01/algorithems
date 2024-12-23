package graph

import (
	"fmt"
	"log"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(key int) {
	if contains(g.vertices, key) {
		log.Fatalf("Vertex %#v exist in the graph ", key)
	} else {
		g.vertices = append(g.vertices, &Vertex{key: key})
	}
}
func (g *Graph) AddEdge(src, des int) {

	srcVert := g.getVertex(src)
	desVert := g.getVertex(des)
	if srcVert == nil || desVert == nil {
		log.Printf("Invalid Edge %#v --> %#v ", src, des)

	} else if contains(srcVert.adjacent, des) {

		log.Printf("Edge exist %#v --> %#v", src, des)
	} else {
		srcVert.adjacent = append(srcVert.adjacent, desVert)
	}

}
func (g *Graph) getVertex(key int) *Vertex {
	for _, k := range g.vertices {
		if k.key == key {
			return g.vertices[key]
		}

	}

	return nil
}
func contains(v []*Vertex, key int) bool {
	for _, k := range v {
		if k.key == key {
			return true
		}
	}

	return false
}

func (g *Graph) Print() {
	for _, vertex := range g.vertices {
		fmt.Printf("\nVertex %v : ", vertex.key)
		for _, v := range vertex.adjacent {
			fmt.Printf("%#v ", v.key)
		}
	}
}
func NewGraph() *Graph {
	g := &Graph{vertices: []*Vertex{}}

	return g
}
