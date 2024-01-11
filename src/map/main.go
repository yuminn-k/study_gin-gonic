package main

import (
	"fmt"
)

func main() {
	// map[key]value
	m := map[string]int{
		"James": 32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)	// map[James:32 Miss Moneypenny:27]
	fmt.Println(m["James"])	// 32
	fmt.Println(m["Barnabas"])	// 0

	v, ok := m["Barnabas"]
	fmt.Println(v)	// 0
	fmt.Println(ok)	// false

	// if m["Barnabas"] exists
	if v, ok := m["Barnabas"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)	// THIS IS THE IF PRINT 27
	}

	// ADD TO MAP
	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}
	/*
		todd 33
		James 32
		Miss Moneypenny 27
	*/
}