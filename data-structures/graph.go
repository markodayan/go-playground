package datastructures

import "fmt"

// Graphs - Adjacency List for Directed Graph

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

func (g *Graph) AddEdge(from int, to int) {
	// get vertex
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	// check for errors (edge already exists, adding edge to vertices that dont exist)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}

}

// Return pointer to specified vertex
func (g *Graph) GetVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}

	return nil
} 

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}

	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, w := range v.adjacent {
			fmt.Printf(" %v", w.key)
		}
	}
	fmt.Println() // Line break for print formatting
}

func TestAdjacencyListGraph() {
	graph := &Graph{}
	
	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}

	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 1)

	graph.AddEdge(6, 2)
	graph.AddEdge(3, 1)
	graph.AddEdge(3, 4)
	
	graph.Print()
}