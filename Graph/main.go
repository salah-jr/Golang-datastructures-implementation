package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if !contains(g.vertices, k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
	} else {
		fmt.Printf("%d already in the list", k)
	}
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}
func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	test.AddEdge(1, 2)
	// try to add duplicate edge
	test.AddEdge(1, 2)
	// try to add invalid edge
	test.AddEdge(6, 2)
	test.AddEdge(0, 3)
	test.AddEdge(0, 4)
	test.AddEdge(2, 3)
	test.AddEdge(4, 1)
	test.AddEdge(1, 3)
	test.AddEdge(1, 4)
	test.AddEdge(3, 4)
	test.AddEdge(3, 0)
	test.Print()

	// fmt.Println()
}
