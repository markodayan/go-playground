package main

import "testing"

func TestKaratsubaMultiplication(t *testing.T) {
	x1 := karatsuba(5678, 1234)
	x1_expected := 7006652
	
	x2 := karatsuba(146123, 352120)
	x2_expected := 51452830760

	if x1 != x1_expected {
		t.Errorf("Expected  %v but got %v", x1_expected, x1)
	}

	if x2 != x2_expected {
		t.Errorf("Expected  %v but got %v", x2_expected, x2)
	}
}