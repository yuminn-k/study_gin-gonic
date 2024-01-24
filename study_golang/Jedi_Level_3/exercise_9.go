/*
	Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER "favSport".
*/

package main

import (
	"fmt"
)

func main() {
	favSport := "basketball"
	switch favSport {
	case "basketball":
		fmt.Println("My favorite sport is basketball")
	case "baseball":
		fmt.Println("My favorite sport is baseball")
	case "football":
		fmt.Println("My favorite sport is football")
	}
}
