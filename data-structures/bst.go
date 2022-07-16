package datastructures

var count int


type node struct {
	val int
	left *node
	right *node
}

func (n *node) Insert(v int) {
 if n.val < v {
	// move right
	if n.right == nil {
		n.right = &node{val:v}
	} else {
		n.right.Insert(v)
	}
 } else if n.val > v {
	// move left
	if n.left == nil {
		n.left = &node{val:v}
	} else {
		n.left.Insert(v)
	}
 }
}

// return true if BST contains node with that value
func (n *node) Search(v int) bool {
	count++

	if n == nil {
		return false
	}

	if n.val < v {
		// move right
		return n.right.Search(v)
	} else if n.val > v {
		// move left
		return n.left.Search(v)
	}

	return true
}

func InitBST(v int) *node {
	return &node{val: v}
}

// func testBST() {
// 	tree := &node{val:5}
// 	tree.Insert(10)
// 	tree.Insert(15)
// 	tree.Insert(3)
// 	fmt.Println(tree)
// 	count = 0
// 	isPresent := tree.Search(7)
// 	fmt.Printf("isPresent: %v . Number of nodes traversed: %v\n", isPresent, count)
// }