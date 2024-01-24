/*
	Building on the previous hands-on exercise, create a program that uses "else if" and "else".
*/

package main

import (
	"fmt"
)

func main() {
	x := "James Bond"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("Bond, James Bond")
	} else {
		fmt.Println("Who are you?")
	}
}
