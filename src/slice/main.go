package main

import (
	"fmt"
)

// a SLICE allows you to group together VALUES of the same TYPE
func main() {
	// COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))	// 5
	fmt.Println(x)	// [4 5 7 8 42]
	fmt.Println(x[0])	// 4
	fmt.Println(x[1])	// 5
	fmt.Println(x[2])	// 7
	fmt.Println(x[3])	// 8
	fmt.Println(x[4])	// 42

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

	fmt.Println(x[:])		// [4 5 7 8 42]
	fmt.Println(x[1:])	// [5 7 8 42]
	fmt.Println(x[1:3])	// [5 7]

	// APPEND
	x = append(x, 77, 88, 99, 1014)	// append values to a slice

	fmt.Println(x)	// [4 5 7 8 42 77 88 99 1014]

	y := []int{234, 456, 678, 987}
	x = append(x, y...)

	fmt.Println(x)	// [4 5 7 8 42 77 88 99 1014 234 456 678 987]

	x = append(x, y)

	fmt.Println(x)	// [4 5 7 8 42 77 88 99 1014 234 456 678 987 [234 456 678 987]]

	// DELETE
	x = append(x[:2], x[4:]...)	// delete values from a slice

	fmt.Println(x)	// [4 5 42 77 88 99 1014 234 456 678 987 [234 456 678 987]]

	// MAKE
	z := make([]int, 10, 100)	// make a slice of length 10 with underlying array of length 100

	fmt.Println(z)	// [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(z))	// 10
	fmt.Println(cap(z))	// 100

	z[0] = 42
	z[9] = 999

	fmt.Println(z)	// [42 0 0 0 0 0 0 0 0 999]
	fmt.Println(len(z))	// 10
	fmt.Println(cap(z))	// 100

	// z[10]	= 333	// index out of range
	z = append(z, 3423)	// append value to slice

	fmt.Println(z)	// [42 0 0 0 0 0 0 0 0 999 3423]
	fmt.Println(len(z))	// 11
	fmt.Println(cap(z))	// 100
}