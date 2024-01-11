package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

func main() {
	p1 := person{
		first: "James",
		last: "Bond",
		age: 32,
	}
	p2 := person{"Miss", "Moneypenny", 27}

	fmt.Println(p1)	// {James Bond 32}
	fmt.Println(p2)	// {Miss Moneypenny 27}

	fmt.Println(p1.first, p1.last, p1.age)	// James Bond 32
}