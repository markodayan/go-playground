package datastructures

import "fmt"

// Stacks are LIFO data structures (Last in first out)

type Stack[T int|string] struct {
	items []T
}

func (s *Stack[T]) Push(v T)  {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() T {
	l := len(s.items)

	toRemove := s.items[l - 1];
	s.items = s.items[:l - 1]
	
	return toRemove
}


func testStack() {
	stack := &Stack[int]{}
	stack.Push(5)
	stack.Push(10)
	stack.Push(15)
	stack.Pop()
	fmt.Println(stack)
}