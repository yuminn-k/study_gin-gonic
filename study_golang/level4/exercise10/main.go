/*
Using the code from the previous example, delete a record from your map.
Now print the map out using the "range" loop
*/
package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james": {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr": {"Being evil", "Ice cream", "Sunsets"},
	}

	delete(x, "no_dr")

	for k, v := range x {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Printf("\tindex position: %v\tvalue: %v\n", i, v2)
		}
	}
}