package main

import "fmt"

// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initialization
var y = 43

// x := 42 // this is not allowed outside of a function
// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nill for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x) // 42

	fmt.Println(y) // 43

	foo()
}

func foo() {
	fmt.Println("again:", y) // again:43
}
