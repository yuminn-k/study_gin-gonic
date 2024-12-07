/*
Create a slice of a slice of string([][]string). Store the following data in the multi-dimensional slice.
- "James", "Bond", "Shaken, not stirred"
- "Miss", "Moneypenny", "Helloooooo, James."
Range over the records, then range over the data in each record.
*/
package main

import "fmt"

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xxs := [][]string{xs1, xs2}
	
	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\tindex position: %v\tvalue: %v\n", j, val)
		}
	}
}
