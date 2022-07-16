package main

import (
	"fmt"
	"playground/algorithms"
	c "playground/crypto"
	datastructures "playground/data-structures"
	"playground/pointers"
	"playground/utils"
)


func main() {
	// a := 8 // 2^3
	// fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 = 64
	// fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 = 1

	// arr := []int{5,2,4,6,1,3, 8, 8 ,1}

	// start := time.Now()
	product := algorithms.Karatsuba(146123, 352120)
	fmt.Println(product)
	// log.Println(product, "Execution Time:", time.Since(start))

	// // x := 146123 / (int(math.Pow10(3)))
	// // fmt.Println(x)

	encoded := c.Base64Encoding("Hello World!")
	fmt.Println("Encoded String:", encoded)

	decoded := c.Base64Decoding(encoded)
	fmt.Println("Decoded String:", decoded)

	sorted := algorithms.InsertionSort([]int{5,4,2,3,1})
	fmt.Println(sorted)

	arr := pointers.IntArr{5,2,4,6,1,3, 8, 8 ,1}
	arr.InsertionSort()
	fmt.Println(arr)

	name,size, err := utils.Capitalize("Bob")

	if err != nil {
		fmt.Println("Could not capitalize", err)
	} else {
		fmt.Printf("Capitalized name: %s, length %d\n", name, size)
	}

	key := c.Key()
	fmt.Println(key)

	woord := "zyzz"
	w := woord[0] - 'a' // 25
	fmt.Println(w)
	

	// testBST()
	// testAdjacencyListGraph()
	// testHashTable()
	// testLinkedList()
	// testMaxHeap()
	// testQueue()
	// testStack()
	testTrie()
}

func testBST() {
	tree := datastructures.InitBST(5)

	toInsert := []int{10, 15, 3}
	for _, v := range toInsert {
		tree.Insert(v)
	}

	fmt.Println(tree)

	isPresent := tree.Search(7)
	fmt.Printf("isPresent: %v\n", isPresent)
}

func testAdjacencyListGraph() {
	graph := &datastructures.Graph{}

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
	graph.AddEdge(4, 1)
	
	graph.Print()
}

func testHashTable() {
	hashTable := datastructures.InitHashTable()

	hashTable.Insert("RANDY")
	hashTable.Insert("ERIC")

	fmt.Println("RANDY search:", hashTable.Search("RANDY"))
	fmt.Println("ALAN search:", hashTable.Search("ALAN"))
	fmt.Println("ERIC search:", hashTable.Search("ERIC"))
	
	hashTable.Delete("RANDY")
	fmt.Println("RANDY search (post delete):", hashTable.Search("RANDY"))
}

func testLinkedList() {
	linkedList := &datastructures.List[int]{}

	linkedList.Push(10)
	linkedList.Push(13)
	linkedList.Push(23)

	fmt.Printf("List: %v\n", linkedList.GetAll())
}

func testMaxHeap() {
	heap := &datastructures.MaxHeap{}

	toInsert := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range toInsert {
		heap.Insert(v)
	}

	for i := 0; i < 5; i++ {
		heap.Extract()
		fmt.Println(heap)
	}
}

func testQueue() {
	queue := &datastructures.Queue[int]{}

	toInsert := []int{5, 10, 15}
	for _,v := range toInsert {
		queue.Enqueue(v)
	}

	fmt.Printf("Removed %v from queue\n", queue.Dequeue())
	fmt.Println(queue)
}

func testStack() {
	stack := &datastructures.Stack[int]{}

	toInsert := []int{5, 10, 15}
	for _, v := range toInsert {
		stack.Push(v)
	}

	fmt.Printf("%v popped off stack\n", stack.Pop())
	fmt.Println(stack)
}

func testTrie() {
	trie := &datastructures.Trie{Root: &datastructures.Node{}}

	toInsert := []string{"orc", "orange"}
	for _, v := range toInsert {
		trie.Insert(v)
	}

	fmt.Printf("or: %v \n", trie.Search("or")) // false
	fmt.Printf("orc: %v \n", trie.Search("orc")) // true
	fmt.Printf("ork: %v \n", trie.Search("ork")) // false
	fmt.Printf("oran: %v \n", trie.Search("oran")) // false
	fmt.Printf("orange: %v \n", trie.Search("orange")) // true
}