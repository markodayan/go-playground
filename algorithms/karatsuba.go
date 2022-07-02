package algorithms

import (
	"math"
	"strconv"
)

// product = 146123 * 352120
// product = ab * cd

// x = a * 10^(n/2) + b -> 146 x 10^3 + 123
// y = c * 10^(n/2) + d -> 352 x 10^3 + 120

// x * y = (a * 10^(n/2) + b)(c * 10^(n/2) + d)
// x * y = ac * 10^(2 * n/2) + (ad + bc) * 10^(n/2) + bd -> we need values of ac, ad + bc, bd

// ac = karatsuba(a,c)
// bd = karatsuba(b,d)

// we can use ac and bd values to determine ad + bc (using Gauss elimination)
// ad + bc = karatsuba(a + b, c + d) - ac - bd

func Karatsuba(x int, y int) int {
	if x < 10 && y < 10 {
		return x * y
	}

	x_length := float64(len(strconv.Itoa(x)))
	y_length := float64(len(strconv.Itoa(y)))

	n := int(math.Max(x_length, y_length))
	half := n / 2

	a := x / (int(math.Pow10(half)))
	c := y / (int(math.Pow10(half)))
	b := x % (int(math.Pow10(half)))
	d := y % (int(math.Pow10(half)))

	ac := Karatsuba(a, c)
	bd := Karatsuba(b, d)
	ad_plus_bc := Karatsuba(a + b, c + d) - ac - bd
	
	return ac * (int(math.Pow10(2 * half))) + (ad_plus_bc * int(math.Pow10(half))) + bd
}