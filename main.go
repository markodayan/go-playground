package main

import (
	"fmt"
	"playground/algorithms"
	c "playground/crypto"
	"playground/pointers"
)


func main() {
	// a := 8 // 2^3
	// fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 = 64
	// fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 = 1

	// arr := []int{5,2,4,6,1,3, 8, 8 ,1}

	// start := time.Now()
	product := algorithms.Karatsuba(146123, 352120)
	fmt.Println(product)
	// log.Println(product, "Execution Time:", time.Since(start))

	// // x := 146123 / (int(math.Pow10(3)))
	// // fmt.Println(x)

	encoded := c.Base64Encoding("Hello World!")
	fmt.Println("Encoded String:", encoded)

	decoded := c.Base64Decoding(encoded)
	fmt.Println("Decoded String:", decoded)

	sorted := algorithms.InsertionSort([]int{5,4,2,3,1})
	fmt.Println(sorted)

	arr := pointers.IntArr{5,2,4,6,1,3, 8, 8 ,1}
	arr.InsertionSort()
	fmt.Println(arr)
}
