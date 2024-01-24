package main

import (
	"fmt"
)

var y int = 42

func main() {
	fmt.Println(y)         // 42
	fmt.Printf("%T\n", y)  // int
	fmt.Printf("%b\n", y)  // 101010, binary
	fmt.Printf("%x\n", y)  // 2a, hex
	fmt.Printf("%#x\n", y) // 0x2a, hex with 0x prefix

	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y)
	fmt.Println(s)      // 0x2a	101010	2a
	fmt.Printf("%v", y) // 42
}
