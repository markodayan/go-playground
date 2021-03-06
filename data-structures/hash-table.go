package datastructures

// Arrays with linked list to manage collisions

import "fmt"

// Need to figure out how to instantiate fixed-size array when initialising hash table
const ArraySize = 7


type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode

}

type bucketNode struct {
	key string
	next *bucketNode
}

// Insert - add to hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// // Search - check if key is stored in hash table 
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// // Delete - remove key from hash table array
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert - add key to linked list
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

// search - check if key is contained by linked list
func (b *bucket) search(k string) bool {
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == k {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

// delete - delete node from linked list
func (b *bucket) delete(k string) {
	// special case: key-to-delete is the head of the linked list
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head

	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete (pointing previous node to node after our specified node to remove)
			previousNode.next = previousNode.next.next
			return
		}

		previousNode = previousNode.next
	}
}

func hash(key string) int {
	total := 0

	for _, c := range key {
		total += int(c)
	}

	return total % ArraySize
}

func InitHashTable() *HashTable {
	result := &HashTable{}
	
	// create linked list for each array index
	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}