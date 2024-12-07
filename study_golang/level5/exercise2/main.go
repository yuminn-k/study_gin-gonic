/*
Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
Access each value in the map. Print out the values, ranging over the slice.
*/
package main

import "fmt"

type person struct {
	first string
	last string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "James",
		last: "Bond",
		favFlavors: []string{
			"Chocolate",
			"Vanilla",
			"Strawberry",
		},
	}

	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		favFlavors: []string{
			"Vanilla",
			"Chocolate",
			"Mint",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
	}
}