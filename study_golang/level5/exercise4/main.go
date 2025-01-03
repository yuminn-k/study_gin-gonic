/*
Create and use an anonymous struct.
*/
package main

import "fmt"

func main() {
	s := struct {
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q": 777,
			"M": 888,
		},
		favDrinks: []string{"Martini", "Water"},
	}

	fmt.Println(s.first)
	
	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
