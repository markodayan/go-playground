package algorithms

import (
	. "sort"
	"testing"
)

type intArr []int

func TestInsertionSort(t *testing.T) {
	a1 := intArr{5,2,4,6,1,3}
	a1_expected := intArr{1,2,3,4,5,6}
	a1_sorted := InsertionSort(a1)

	a2 := intArr{2,1,9,76,4}
	a2_expected := intArr{1,2,4,9,76}
	a2_sorted := InsertionSort(a2)

	a3 := intArr{5,3,4,6,2}
	a3_expected := intArr{2,3,4,5,6}
	a3_sorted := InsertionSort(a3)

	a4 := intArr{5,3,4,1,2}
	a4_expected := intArr{1,2,3,4,5}
	a4_sorted := InsertionSort(a4)

	a5 := intArr{10,80,30,90,40,50,70}
	a5_expected := intArr{10,30,40,50,70,80,90}
	a5_sorted := InsertionSort(a5)

	a6 := intArr{1,2,3,4,5,6}
	a6_expected := intArr{1,2,3,4,5,6}
	a6_sorted := InsertionSort(a6)

	if !IntsAreSorted(a1_sorted) {
		t.Errorf("Expected slice %v but got %v", a1_expected, a1_sorted)
	}

	if !IntsAreSorted(a2_sorted) {
		t.Errorf("Expected slice %v but got %v", a2_expected, a2_sorted)
	}

	if !IntsAreSorted(a3_sorted) {
		t.Errorf("Expected slice %v but got %v", a3_expected, a3_sorted)
	}

	if !IntsAreSorted(a4_sorted) {
		t.Errorf("Expected slice %v but got %v", a4_expected, a4_sorted)
	}

	if !IntsAreSorted(a5_sorted) {
		t.Errorf("Expected slice %v but got %v", a5_expected, a5_sorted)
	}

	if !IntsAreSorted(a6_sorted) {
		t.Errorf("Expected slice %v but got %v", a6_expected, a6_sorted)
	}
}

func TestBubbleSort(t *testing.T) {
	a1 := intArr{5,2,4,6,1,3}
	a1_expected := intArr{1,2,3,4,5,6}
	a1_sorted := BubbleSort(a1)

	a2 := intArr{2,1,9,76,4}
	a2_expected := intArr{1,2,4,9,76}
	a2_sorted := BubbleSort(a2)

	a3 := intArr{5,3,4,6,2}
	a3_expected := intArr{2,3,4,5,6}
	a3_sorted := BubbleSort(a3)

	a4 := intArr{5,3,4,1,2}
	a4_expected := intArr{1,2,3,4,5}
	a4_sorted := BubbleSort(a4)

	a5 := intArr{10,80,30,90,40,50,70}
	a5_expected := intArr{10,30,40,50,70,80,90}
	a5_sorted := BubbleSort(a5)

	a6 := intArr{1,2,3,4,5,6}
	a6_expected := intArr{1,2,3,4,5,6}
	a6_sorted := BubbleSort(a6)

	if !IntsAreSorted(a1_sorted) {
		t.Errorf("Expected slice %v but got %v", a1_expected, a1_sorted)
	}

	if !IntsAreSorted(a2_sorted) {
		t.Errorf("Expected slice %v but got %v", a2_expected, a2_sorted)
	}

	if !IntsAreSorted(a3_sorted) {
		t.Errorf("Expected slice %v but got %v", a3_expected, a3_sorted)
	}

	if !IntsAreSorted(a4_sorted) {
		t.Errorf("Expected slice %v but got %v", a4_expected, a4_sorted)
	}

	if !IntsAreSorted(a5_sorted) {
		t.Errorf("Expected slice %v but got %v", a5_expected, a5_sorted)
	}

	if !IntsAreSorted(a6_sorted) {
		t.Errorf("Expected slice %v but got %v", a6_expected, a6_sorted)
	}
}

func TestSelectionSort(t *testing.T) {
	a1 := intArr{5,2,4,6,1,3}
	a1_expected := intArr{1,2,3,4,5,6}
	a1_sorted := SelectionSort(a1)
	
	a2 := intArr{2,1,9,76,4}
	a2_expected := intArr{1,2,4,9,76}
	a2_sorted := SelectionSort(a2)
	
	a3 := intArr{5,3,4,6,2}
	a3_expected := intArr{2,3,4,5,6}
	a3_sorted := SelectionSort(a3)
	
	a4 := intArr{5,3,4,1,2}
	a4_expected := intArr{1,2,3,4,5}
	a4_sorted := SelectionSort(a4)
	
	a5 := intArr{10,80,30,90,40,50,70}
	a5_expected := intArr{10,30,40,50,70,80,90}
	a5_sorted := SelectionSort(a5)
	
	a6 := intArr{1,2,3,4,5,6}
	a6_expected := intArr{1,2,3,4,5,6}
	a6_sorted := SelectionSort(a6)



	if !IntsAreSorted(a1_sorted) {
		t.Errorf("Expected slice %v but got %v", a1_expected, a1_sorted)
	}

	if !IntsAreSorted(a2_sorted) {
		t.Errorf("Expected slice %v but got %v", a2_expected, a2_sorted)
	}

	if !IntsAreSorted(a3_sorted) {
		t.Errorf("Expected slice %v but got %v", a3_expected, a3_sorted)
	}

	if !IntsAreSorted(a4_sorted) {
		t.Errorf("Expected slice %v but got %v", a4_expected, a4_sorted)
	}

	if !IntsAreSorted(a5_sorted) {
		t.Errorf("Expected slice %v but got %v", a5_expected, a5_sorted)
	}

	if !IntsAreSorted(a6_sorted) {
		t.Errorf("Expected slice %v but got %v", a6_expected, a6_sorted)
	}
}

func TestMerge(t *testing.T) {
	a1_sorted := merge([]int{2,5}, []int{4,6})
	a1_expected := []int{2,4,5,6}

	a2_sorted := merge([]int{1,2,5}, []int{4})
	a2_expected := []int{1,2,4,5}

	if !IntsAreSorted(a1_sorted) {
		t.Errorf("Expected slice %v but got %v", a1_expected, a1_sorted)
	}

	if !IntsAreSorted(a2_sorted) {
		t.Errorf("Expected slice %v but got %v", a2_expected, a2_sorted)
	}
}

func TestMergeSort(t *testing.T) {
	a1 := intArr{5,2,4,6,1,3}
	a1_expected := intArr{1,2,3,4,5,6}
	a1_sorted := MergeSort(a1)
	
	a2 := intArr{2,1,9,76,4}
	a2_expected := intArr{1,2,4,9,76}
	a2_sorted := MergeSort(a2)
	
	a3 := intArr{5,3,4,6,2}
	a3_expected := intArr{2,3,4,5,6}
	a3_sorted := MergeSort(a3)
	
	a4 := intArr{5,3,4,1,2}
	a4_expected := intArr{1,2,3,4,5}
	a4_sorted := MergeSort(a4)
	
	a5 := intArr{10,80,30,90,40,50,70}
	a5_expected := intArr{10,30,40,50,70,80,90}
	a5_sorted := MergeSort(a5)
	
	a6 := intArr{1,2,3,4,5,6}
	a6_expected := intArr{1,2,3,4,5,6}
	a6_sorted := MergeSort(a6)



	if !IntsAreSorted(a1_sorted) {
		t.Errorf("Expected slice %v but got %v", a1_expected, a1_sorted)
	}

	if !IntsAreSorted(a2_sorted) {
		t.Errorf("Expected slice %v but got %v", a2_expected, a2_sorted)
	}

	if !IntsAreSorted(a3_sorted) {
		t.Errorf("Expected slice %v but got %v", a3_expected, a3_sorted)
	}

	if !IntsAreSorted(a4_sorted) {
		t.Errorf("Expected slice %v but got %v", a4_expected, a4_sorted)
	}

	if !IntsAreSorted(a5_sorted) {
		t.Errorf("Expected slice %v but got %v", a5_expected, a5_sorted)
	}

	if !IntsAreSorted(a6_sorted) {
		t.Errorf("Expected slice %v but got %v", a6_expected, a6_sorted)
	}
}