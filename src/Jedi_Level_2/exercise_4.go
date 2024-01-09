/*
	Write a program that
		- assigns an int to a variable
		- prints that int in decimal, binary, and hex
		- shifts the bits of that int over i position to the left, and assigns that to a variable
		- prints that variable in decimal, binary, and hex
*/

package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)	// 42	101010	0x2a

	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)	// 84	1010100	0x54
}