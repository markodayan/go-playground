package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSwap(t *testing.T) {
	arr := intArr{1,2,3,4,5,6}
	arr.swap(0, 1)

	if !reflect.DeepEqual(arr, intArr{2,1,3,4,5,6}) {
		t.Errorf("Expected slice %v but got %v", intArr{2,1,3,4,5,6}, arr)
	}
}

func TestMutativeInsertionSort(t *testing.T) {
	a1 := intArr{5,2,4,6,1,3}
	a1_expected := intArr{1,2,3,4,5,6}

	a2 := intArr{2,1,9,76,4}
	a2_expected := intArr{1,2,4,9,76}
	
	a3 := intArr{5,3,4,6,2}
	a3_expected := intArr{2,3,4,5,6}
	
	a4 := intArr{5,3,4,1,2}
	a4_expected := intArr{1,2,3,4,5}
	
	a5 := intArr{10,80,30,90,40,50,70}
	a5_expected := intArr{10,30,40,50,70,80,90}

	a6 := intArr{1,2,3,4,5,6}
	a6_expected := intArr{1,2,3,4,5,6}

	a1.mutativeInsertionSort()
	a2.mutativeInsertionSort()
	a3.mutativeInsertionSort()
	a4.mutativeInsertionSort()
	a5.mutativeInsertionSort()
	a6.mutativeInsertionSort()

	if !reflect.DeepEqual(a1, a1_expected) {
		t.Errorf("Expected slice %v but got %v", a1_expected, a1)
	}

	if !reflect.DeepEqual(a2, a2_expected) {
		t.Errorf("Expected slice %v but got %v", a2_expected, a2)
	}

	if !reflect.DeepEqual(a3, a3_expected) {
		t.Errorf("Expected slice %v but got %v", a3_expected, a3)
	}

	if !reflect.DeepEqual(a4, a4_expected) {
		t.Errorf("Expected slice %v but got %v", a4_expected, a4)
	}

	if !reflect.DeepEqual(a5, a5_expected) {
		t.Errorf("Expected slice %v but got %v", a5_expected, a5)
	}

	if !reflect.DeepEqual(a6, a6_expected) {
		t.Errorf("Expected slice %v but got %v", a6_expected, a6)
	}
}

func TestMutativeBubbleSort(t *testing.T) {
	a1 := intArr{5,2,4,6,1,3}
	a1_expected := intArr{1,2,3,4,5,6}

	a2 := intArr{2,1,9,76,4}
	a2_expected := intArr{1,2,4,9,76}
	
	a3 := intArr{5,3,4,6,2}
	a3_expected := intArr{2,3,4,5,6}
	
	a4 := intArr{5,3,4,1,2}
	a4_expected := intArr{1,2,3,4,5}
	
	a5 := intArr{10,80,30,90,40,50,70}
	a5_expected := intArr{10,30,40,50,70,80,90}

	a6 := intArr{1,2,3,4,5,6}
	a6_expected := intArr{1,2,3,4,5,6}

	a1.mutativeBubbleSort()
	a2.mutativeBubbleSort()
	a3.mutativeBubbleSort()
	a4.mutativeBubbleSort()
	a5.mutativeBubbleSort()
	a6.mutativeBubbleSort()

	if !reflect.DeepEqual(a1, a1_expected) {
		t.Errorf("Expected slice %v but got %v", a1_expected, a1)
	}

	if !reflect.DeepEqual(a2, a2_expected) {
		t.Errorf("Expected slice %v but got %v", a2_expected, a2)
	}

	if !reflect.DeepEqual(a3, a3_expected) {
		t.Errorf("Expected slice %v but got %v", a3_expected, a3)
	}

	if !reflect.DeepEqual(a4, a4_expected) {
		t.Errorf("Expected slice %v but got %v", a4_expected, a4)
	}

	if !reflect.DeepEqual(a5, a5_expected) {
		t.Errorf("Expected slice %v but got %v", a5_expected, a5)
	}

	if !reflect.DeepEqual(a6, a6_expected) {
		t.Errorf("Expected slice %v but got %v", a6_expected, a6)
	}
}

func TestMutativeQuickSort(t *testing.T) {
	a1 := intArr{5,2,4,6,1,3}
	a1_expected := intArr{1,2,3,4,5,6}

	a2 := intArr{2,1,9,76,4}
	a2_expected := intArr{1,2,4,9,76}
	
	a3 := intArr{5,3,4,6,2}
	a3_expected := intArr{2,3,4,5,6}
	
	a4 := intArr{5,3,4,1,2}
	a4_expected := intArr{1,2,3,4,5}
	
	a5 := intArr{10,80,30,90,40,50,70}
	a5_expected := intArr{10,30,40,50,70,80,90}

	a6 := intArr{1,2,3,4,5,6}
	a6_expected := intArr{1,2,3,4,5,6}

	a1.mutativeQuickSort(0, len(a1) - 1)
	a2.mutativeQuickSort(0, len(a2) - 1)
	a3.mutativeQuickSort(0, len(a3) - 1)
	a4.mutativeQuickSort(0, len(a4) - 1)
	a5.mutativeQuickSort(0, len(a5) - 1)
	a6.mutativeQuickSort(0, len(a6) - 1)

	if !reflect.DeepEqual(a1, a1_expected) {
		t.Errorf("Expected slice %v but got %v", a1_expected, a1)
	}

	if !reflect.DeepEqual(a2, a2_expected) {
		t.Errorf("Expected slice %v but got %v", a2_expected, a2)
	}

	if !reflect.DeepEqual(a3, a3_expected) {
		t.Errorf("Expected slice %v but got %v", a3_expected, a3)
	}

	if !reflect.DeepEqual(a4, a4_expected) {
		t.Errorf("Expected slice %v but got %v", a4_expected, a4)
	}

	if !sort.IntsAreSorted(a5)  {
		t.Errorf("Expected slice %v but got %v", a5_expected, a5)
	}

	if !sort.IntsAreSorted(a6) {
		t.Errorf("Expected slice %v but got %v", a6_expected, a6)
	}
}