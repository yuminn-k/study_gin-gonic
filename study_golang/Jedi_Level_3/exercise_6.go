/*
	Create a program that shows the "if statement" in action.
*/

package main

import (
	"fmt"
)

func main() {
	x := "James Bond"

	if x == "James Bond" {
		fmt.Println("Bond, James Bond")
	} else {
		fmt.Println("This is not James Bond")
	}
}
