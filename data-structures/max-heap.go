package datastructures

import "fmt"

// Heapify = the process of rearranging indexes (we typically heapify after inserting or deleting elements)

type MaxHeap struct{
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	// Move key to correct position in the max heap
	h.maxHeapifyUp(len(h.array) - 1)
}

// Return largest key, and remove it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	last := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	// Get last key and put it in the root index (this is how we start)
	h.array[0] = h.array[last]
	h.array = h.array[:last] // slice last element off (which is now the first element)

	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
 for h.array[parent(index)] < h.array[index] {
 		h.swap(parent(index), index)
	 	index  = parent(index)
 }
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has atleast one child
	for l <= lastIndex {

		// when left child is only child
		if l == lastIndex {
			childToCompare = l
			} else if h.array[l] > h.array[r] {
			// when left child is larger
				childToCompare = l
		} else {
			// when right child is larger
			childToCompare = r
		}

		// compare array val of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) swap(i1,i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return (i * 2) + 1
}

func right(i int) int {
	return (i * 2) + 2 
}

func TestHeap() {
	maxHeap := &MaxHeap{}
	
	buildHeap := []int{10,20,30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		maxHeap.Insert(v)
		// fmt.Println(maxHeap)
	}
	
	for i := 0; i < 5; i++ {
		maxHeap.Extract()
		fmt.Println(maxHeap)
	}


}