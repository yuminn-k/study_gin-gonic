/*
Create a map with a key of TYPE string which is a person's last_first name, and a value of type []string which stores their favorite things.
Store three records in your map.
Print out all of the values, along with their index position in the slice, without using the range clause.
*/
package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james": {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr": {"Being evil", "Ice cream", "Sunsets"},
	}

	for k, v := range x {
		fmt.Println("This is the record for ", k)
		for i, val := range v {
			fmt.Printf("\tindex position: %v\tvalue: %v\n", i, val)
		}
	}
}
