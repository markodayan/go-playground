package datastructures

import "testing"

func TestBST(t *testing.T) {
	tree := InitBST(5)

	toInsert := []int{10, 15, 3}
	for _, v := range toInsert {
		tree.Insert(v)
	}

	notInTree := []int{7, 8, 2}

	for _, v := range notInTree {
		isPresent := tree.Search(v)

		if isPresent == true {
			t.Errorf("[Expected value not in tree] - tree.Search(%v) should be false but got %t", v, isPresent)
		}
	}

	for _, v := range toInsert {
		isPresent := tree.Search(v)

		if isPresent == false {
			t.Errorf("[Expected value in tree] - tree.Search(%v) should be true but got %t", v, isPresent)
		}
	}
}

func TestAdjacencyListGraph(t *testing.T) {
	graph := &Graph{}

	// Add valid vertices
	for i := 0; i < 5; i++ {
		err := graph.AddVertex(i)

		if err != nil {
			t.Errorf("[AddVertex()] - %s", err.Error())
		}
	}

	type EdgeTuple struct {
		from int
		to int
	}
	
	type Edges []EdgeTuple

	edges := Edges{
		EdgeTuple{0,3}, 
		EdgeTuple{1,2},
		EdgeTuple{1,3},
		EdgeTuple{1,4},
		EdgeTuple{2,3},
		EdgeTuple{3,1},
	}

	// Add valid edges
	for _, v := range edges {
		err := graph.AddEdge(v.from, v.to)
		if err != nil {
			t.Errorf("[AddEdge()] - %s", err.Error())
		}
	}

	// Expect failure for adding vertices that already exist
	existingVertices := []int{0, 4}
	for _, v := range existingVertices {
		err := graph.AddVertex(v)
		if err == nil {
			t.Errorf("Expected error for AddVertex(%v) but got none", v)
		}
	}

	// Expect failure for attempting to add existing edge
	existingEdges := Edges{
		EdgeTuple{0,3},
		EdgeTuple{2,3},
		EdgeTuple{1,2},
	}

	for _, v := range existingEdges {
		err := graph.AddEdge(v.from, v.to)

		if err == nil {
			t.Errorf("Expected error (attempting to add existing edge) for AddEdge(%v, %v) but got none", v.from, v.to)
		}
	}

	// Expect failure for attempting to add invalid edge
	invalidEdges := Edges{
		EdgeTuple{5, 6},
		EdgeTuple{1, 100},
		EdgeTuple{5, 10},
		EdgeTuple{4, -1},
	}

	for _, v := range invalidEdges {
		err := graph.AddEdge(v.from, v.to)

		if err == nil {
			t.Errorf("Expected error (attempting to add invalid edge) for AddEdge(%v, %v) but got none", v.from, v.to)
		}
	}
}

func TestHashTable(t *testing.T) {
	ht := InitHashTable()

	validKeys := []string{"RANDY", "Eric"}
	invalidKeys := []string{"BOB", "ALICE"}

	for _, v := range validKeys {
		ht.Insert(v)
		containsKey := ht.Search(v)
	
		if containsKey == false {
			t.Errorf("[Expected key in hash table] - ht.Search(%v) should be true but got %t", v, containsKey)
		}
	}

	for _, v := range invalidKeys {
		containsKey := ht.Search(v)

		if containsKey == true {
			t.Errorf("[Expected missing key in hash table] - ht.Search(%v) should be false but got %t", v, containsKey)
		}
	}

	ht.Delete("RANDY")
	containsKey := ht.Search("RANDY")

	if containsKey == true {
		t.Errorf("[Expected missing key in hash table] - ht.Search(%v) should be false but got %t", "RANDY", containsKey)
	}
}
