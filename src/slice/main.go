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
}