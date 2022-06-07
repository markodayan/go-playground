package main

import (
	"fmt"
)
type intArr []int

func main() {
	// a := 8 // 2^3
	// fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 = 64
	// fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 = 1

	// arr := []int{5,2,4,6,1,3, 8, 8 ,1}

	product := karatsuba(146123, 352120)
	fmt.Println(product)

	// x := 146123 / (int(math.Pow10(3)))
	// fmt.Println(x)

	encoded := base64Encoding("Hello World!")
	fmt.Println("Encoded String:", encoded)

	decoded := base64Decoding(encoded)
	fmt.Println("Decoded String:", decoded)
}
