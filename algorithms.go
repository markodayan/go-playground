package main

func (arr *intArr) swap(i1 int, i2 int) {
	(*arr)[i1], (*arr)[i2] = (*arr)[i2], (*arr)[i1]
}

func (arrPointer *intArr) partition(low int, high int) int {
	// last array element selected as pivot
	pivot := (*arrPointer)[high]
	i := low - 1

	for j := low; j <=high - 1; j++ {
		if (*arrPointer)[j] < pivot {
			i++;
			(*arrPointer).swap(i, j)
		}
	}
	
	(*arrPointer).swap(i + 1, high)
	return i + 1;
}

func (arrPointer *intArr) insertionSort() {
	for i := 1; i < len(*arrPointer); i++ {
		for j := i - 1; j >= 0; j-- {
			if (*arrPointer)[j] > (*arrPointer)[j + 1] {
				(*arrPointer).swap(j, j + 1)
			}
		} 
	}                                                           
}

func (arrPointer *intArr) bubbleSort() {
	var noSwaps bool

	for i := 0; i < len(*arrPointer) - 1; i++ {
		noSwaps = true 	 
		for j := 0; j < len(*arrPointer) - 1 - i; j++ {
			if (*arrPointer)[j] > (*arrPointer)[j + 1] {
				(*arrPointer).swap(j, j + 1)
				noSwaps = false
			}
		}

		if noSwaps {
			break
		}
	}
}

func (arrPointer *intArr) quickSort(low int, high int) {
	if low < high {
		partitionIndex := (*arrPointer).partition(low, high)
		(*arrPointer).quickSort(low, partitionIndex - 1)
		(*arrPointer).quickSort(partitionIndex + 1, high)
	}
}