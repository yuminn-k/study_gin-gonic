package main

import (
	"fmt"
)

var x bool

func main() {
	a := 7
	b := 42
	fmt.Println(a == b) // false
	fmt.Println(a <= b) // true
	fmt.Println(x)      // false
	x = true
	fmt.Println(x) // true
}
