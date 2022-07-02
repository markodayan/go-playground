package pointers

func (arrPointer *IntArr) push(val int) {
	*arrPointer = append(*arrPointer, val)
}