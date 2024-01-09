/*
	Create TYPED and UNTYPED constants. Print the values of the constants
*/

package main

import (
	"fmt"
)

const (
	a = 42			// untyped constant
	b int = 42	// typed constant
)

func main() {
	fmt.Println(a, b)	// 42 42
}