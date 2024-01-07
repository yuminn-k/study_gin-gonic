package main

import (
	"fmt"
)

var y string
var z int

func main() {
	// DECLARE a variable to be of a certain TYPE
	// and then ASSIGN a VALUE of that type to the variable

	fmt.Println("printing the value of y", y, "ending")	// printing the value of y  ending
	fmt.Printf("%T\n", y)	// string
	
	y = "Bond, James Bond"

	fmt.Println("printing the value of y", y, "ending")	// printing the value of y Bond, James Bond ending
	fmt.Printf("%T\n", y)	// string

	fmt.Println("printing the value of z", z, "ending")	// printing the value of z 0 ending
	fmt.Printf("%T", z)	// int
}