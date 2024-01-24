/*
	Using iota, create 4 constants for the last 4 years. Print the constant values.
*/

package main

import (
	"fmt"
)

const (
	a = iota + 2017
	b = iota + 2017
	c = iota + 2017
	d = iota + 2017
)

func main() {
	fmt.Println(a) // 2017
	fmt.Println(b) // 2018
	fmt.Println(c) // 2019
	fmt.Println(d) // 2020
}
