package main

import (
	"fmt"
)

var y = 42

// DECLARE the VARiABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "shaken, not stirred"
var z string = "Shaken, not stirred"
var a string = `James said, "Shaken, not stirred"`

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

func main() {
	fmt.Println(y)	// 42
	fmt.Printf("%T\n", y)	// int
	fmt.Println(z)	// Shaken, not stirred
	fmt.Printf("%T\n", z)	// string
	fmt.Println(a)	// James said, "Shaken, not stirred"
	fmt.Printf("%T\n", a)	// string
}