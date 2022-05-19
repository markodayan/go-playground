package main

import (
	"fmt"
	"reflect"
	"strconv"
)

const CURRENT_MONTH string = "December"
const CURRENT_YEAR = 2022 

func main() {
	var (
		i = 10
		s = "Hello"
	)

	fmt.Println(CURRENT_MONTH)
	fmt.Println(CURRENT_YEAR)

	fmt.Println(i, reflect.TypeOf(i))
	fmt.Println(s, reflect.TypeOf(s))

	// Type Conversions
	// string to integer conversion
	intVar1, err := strconv.Atoi("100")
	fmt.Println(intVar1, err, reflect.TypeOf(intVar1)) // 100 <nil> int

	// string to integer (specify base)
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	strVar := "100"
	intVar, err := strconv.ParseInt(strVar, 2, 8)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 64)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}