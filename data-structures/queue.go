package datastructures

import "fmt"

// Queues are FIFO data structures (First in first out)

type Queue[T int|string] struct {
	items []T
}


func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Dequeue() T {
	toRemove := q.items[0]
	q.items = q.items[1:]

	return toRemove
}

func testQueue() {
	queue := &Queue[int]{}
	queue.Enqueue(5)
	queue.Enqueue(10)
	queue.Enqueue(15)
	removed := queue.Dequeue()
	fmt.Println(removed)
	fmt.Println(queue)
}