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
}