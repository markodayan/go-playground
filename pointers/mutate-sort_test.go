package pointers

import (
	"reflect"
	"sort"
	"testing"
)


func TestSwap(t *testing.T) {
	arr := IntArr{1,2,3,4,5,6}
	arr.swap(0, 1)

	if !reflect.DeepEqual(arr, IntArr{2,1,3,4,5,6}) {
		t.Errorf("Expected slice %v but got %v", IntArr{2,1,3,4,5,6}, arr)
	}
}

func TestInsertionSort(t *testing.T) {
	a1 := IntArr{5,2,4,6,1,3}
	a1_expected := IntArr{1,2,3,4,5,6}

	a2 := IntArr{2,1,9,76,4}
	a2_expected := IntArr{1,2,4,9,76}
	
	a3 := IntArr{5,3,4,6,2}
	a3_expected := IntArr{2,3,4,5,6}
	
	a4 := IntArr{5,3,4,1,2}
	a4_expected := IntArr{1,2,3,4,5}
	
	a5 := IntArr{10,80,30,90,40,50,70}
	a5_expected := IntArr{10,30,40,50,70,80,90}

	a6 := IntArr{1,2,3,4,5,6}
	a6_expected := IntArr{1,2,3,4,5,6}

	a1.InsertionSort()
	a2.InsertionSort()
	a3.InsertionSort()
	a4.InsertionSort()
	a5.InsertionSort()
	a6.InsertionSort()

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

func TestBubbleSort(t *testing.T) {
	a1 := IntArr{5,2,4,6,1,3}
	a1_expected := IntArr{1,2,3,4,5,6}

	a2 := IntArr{2,1,9,76,4}
	a2_expected := IntArr{1,2,4,9,76}
	
	a3 := IntArr{5,3,4,6,2}
	a3_expected := IntArr{2,3,4,5,6}
	
	a4 := IntArr{5,3,4,1,2}
	a4_expected := IntArr{1,2,3,4,5}
	
	a5 := IntArr{10,80,30,90,40,50,70}
	a5_expected := IntArr{10,30,40,50,70,80,90}

	a6 := IntArr{1,2,3,4,5,6}
	a6_expected := IntArr{1,2,3,4,5,6}

	a1.BubbleSort()
	a2.BubbleSort()
	a3.BubbleSort()
	a4.BubbleSort()
	a5.BubbleSort()
	a6.BubbleSort()

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

func TestQuickSort(t *testing.T) {
	a1 := IntArr{5,2,4,6,1,3}
	a1_expected := IntArr{1,2,3,4,5,6}

	a2 := IntArr{2,1,9,76,4}
	a2_expected := IntArr{1,2,4,9,76}
	
	a3 := IntArr{5,3,4,6,2}
	a3_expected := IntArr{2,3,4,5,6}
	
	a4 := IntArr{5,3,4,1,2}
	a4_expected := IntArr{1,2,3,4,5}
	
	a5 := IntArr{10,80,30,90,40,50,70}
	a5_expected := IntArr{10,30,40,50,70,80,90}

	a6 := IntArr{1,2,3,4,5,6}
	a6_expected := IntArr{1,2,3,4,5,6}

	a1.QuickSort(0, len(a1) - 1)
	a2.QuickSort(0, len(a2) - 1)
	a3.QuickSort(0, len(a3) - 1)
	a4.QuickSort(0, len(a4) - 1)
	a5.QuickSort(0, len(a5) - 1)
	a6.QuickSort(0, len(a6) - 1)

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