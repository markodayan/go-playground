package main

import (
	. "sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a1 := intArr{5,2,4,6,1,3}
	a1_expected := intArr{1,2,3,4,5,6}
	a1_sorted := insertionSort(a1)

	a2 := intArr{2,1,9,76,4}
	a2_expected := intArr{1,2,4,9,76}
	a2_sorted := insertionSort(a2)

	a3 := intArr{5,3,4,6,2}
	a3_expected := intArr{2,3,4,5,6}
	a3_sorted := insertionSort(a3)

	a4 := intArr{5,3,4,1,2}
	a4_expected := intArr{1,2,3,4,5}
	a4_sorted := insertionSort(a4)

	a5 := intArr{10,80,30,90,40,50,70}
	a5_expected := intArr{10,30,40,50,70,80,90}
	a5_sorted := insertionSort(a5)

	a6 := intArr{1,2,3,4,5,6}
	a6_expected := intArr{1,2,3,4,5,6}
	a6_sorted := insertionSort(a6)

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
	a1_sorted := bubbleSort(a1)

	a2 := intArr{2,1,9,76,4}
	a2_expected := intArr{1,2,4,9,76}
	a2_sorted := bubbleSort(a2)

	a3 := intArr{5,3,4,6,2}
	a3_expected := intArr{2,3,4,5,6}
	a3_sorted := bubbleSort(a3)

	a4 := intArr{5,3,4,1,2}
	a4_expected := intArr{1,2,3,4,5}
	a4_sorted := bubbleSort(a4)

	a5 := intArr{10,80,30,90,40,50,70}
	a5_expected := intArr{10,30,40,50,70,80,90}
	a5_sorted := bubbleSort(a5)

	a6 := intArr{1,2,3,4,5,6}
	a6_expected := intArr{1,2,3,4,5,6}
	a6_sorted := bubbleSort(a6)

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
	a1_sorted := selectionSort(a1)
	
	a2 := intArr{2,1,9,76,4}
	a2_expected := intArr{1,2,4,9,76}
	a2_sorted := selectionSort(a2)
	
	a3 := intArr{5,3,4,6,2}
	a3_expected := intArr{2,3,4,5,6}
	a3_sorted := selectionSort(a3)
	
	a4 := intArr{5,3,4,1,2}
	a4_expected := intArr{1,2,3,4,5}
	a4_sorted := selectionSort(a4)
	
	a5 := intArr{10,80,30,90,40,50,70}
	a5_expected := intArr{10,30,40,50,70,80,90}
	a5_sorted := selectionSort(a5)
	
	a6 := intArr{1,2,3,4,5,6}
	a6_expected := intArr{1,2,3,4,5,6}
	a6_sorted := selectionSort(a6)



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
	a1_sorted := mergeSort(a1)
	
	a2 := intArr{2,1,9,76,4}
	a2_expected := intArr{1,2,4,9,76}
	a2_sorted := mergeSort(a2)
	
	a3 := intArr{5,3,4,6,2}
	a3_expected := intArr{2,3,4,5,6}
	a3_sorted := mergeSort(a3)
	
	a4 := intArr{5,3,4,1,2}
	a4_expected := intArr{1,2,3,4,5}
	a4_sorted := mergeSort(a4)
	
	a5 := intArr{10,80,30,90,40,50,70}
	a5_expected := intArr{10,30,40,50,70,80,90}
	a5_sorted := mergeSort(a5)
	
	a6 := intArr{1,2,3,4,5,6}
	a6_expected := intArr{1,2,3,4,5,6}
	a6_sorted := mergeSort(a6)



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