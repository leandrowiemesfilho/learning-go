package main

import "fmt"

/*
Using the previous exercise, add an entry to the map and demonstrate the entire map using range.
*/

func main() {
	m := map[string][]string{
		"LeBron_James": {
			"Play Basketball",
			"Beat records",
		},
		"Jordan_Michael": {
			"Play Basketball",
			"Be the G.O.A.T",
			"Make amazing shoes",
		},
		"Rose_Derrick": {
			"Play Basketball",
			"Make standing dunks",
			"Have an explosive game style",
		},
	}

	m["Durant_Kevin"] = []string{
		"Play Basketball",
		"Beat records",
	}

	for key, value := range m {
		fmt.Println(key)
		for _, v := range value {
			fmt.Printf("\t%v\n", v)
		}
	}
}
