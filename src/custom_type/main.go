package main

import (
	"fmt"
)

var a int
type hotdog int	// custom type
var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)	// 42
	fmt.Printf("%T\n", a)	// int
	fmt.Println(b)	// 43
	fmt.Printf("%T\n", b)	// main.hotdog
	// a = b	// error. different types
}