package algorithms

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}

	return arr
}

	/* Loop over sequence and move largest element to the end */
	func BubbleSort(arr []int) []int {
		/* we use i to control the amount of inner loops after each bubble */
		for i := 0; i < len(arr); i++ {
			noSwaps := true

			for j := 0; j < len(arr) - i - 1; j++ {
				if arr[j] > arr[j + 1] {
					arr[j], arr[j + 1] = arr[j + 1], arr[j]
					noSwaps = false
				}
			}

			if noSwaps {
				break
			}
		}

		return arr
	}


	/* Similar to bubble sort but we move smallest element to start */
	func SelectionSort(arr []int) []int {
		for i := 0; i < len(arr); i++ {
			candidateUpdated := false
			candidateIndex := i

			for j := i + 1; j < len(arr); j++ {
				if arr[j] < arr[candidateIndex] {
					candidateIndex = j
					candidateUpdated = true
				}
			}

			if candidateUpdated {
				arr[i], arr[candidateIndex] = arr[candidateIndex], arr[i]
			}

			
		}		

		return arr
	}

	func MergeSort(arr []int) []int {
		if len(arr) < 2 {
			return arr
		}

		mid := len(arr) / 2
		left := MergeSort(arr[:mid])
		right := MergeSort(arr[mid:])

		return merge(left, right)
	}


	func merge(arr1 []int, arr2 []int) []int {
		i := 0
		j := 0
		resultIndex := 0
		result := make([]int, len(arr1) + len(arr2))

		for i < len(arr1) && j < len(arr2) {
			if arr1[i] < arr2[j] {
				result[resultIndex] = arr1[i]
				i++
			} else {
				result[resultIndex] = arr2[j]
				j++
			}

			resultIndex++
		}

		for i < len(arr1) {
			result[resultIndex] = arr1[i]
			i++
			resultIndex++
		}

		for j < len(arr2) {
			result[resultIndex] = arr2[j]
			j++
			resultIndex++
		}

		return result
	}