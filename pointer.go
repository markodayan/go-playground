package main

func (arrPointer *intArr) push(val int) {
	*arrPointer = append(*arrPointer, val)
}