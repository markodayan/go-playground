package datastructures

// MaxHeap struct has a slice that holds the array
type MaxHeap struct{
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap

// mmaxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
 for h.array[parent(index)] < h.array[index] {
	 h.swap(parent(index), index)
	 index  = parent(index)
 }
}

// Get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// Get the left child index 
func left(i int) int {
	return (i * 2) + 1
}

// Get the right child index
func right(i int) int {
	return (i * 2) + 2 
}

// swap keys in the array
func (h *MaxHeap) swap(i1,i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}