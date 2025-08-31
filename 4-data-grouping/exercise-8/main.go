package main

import "fmt"

/*
Create a map with a string key and a slice of string value.
Key should contain names in the format last_first.
Value should contain the person's favorite hobbies.
Show all these values and their indexes.
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

	for key, value := range m {
		fmt.Println(key)
		for _, v := range value {
			fmt.Printf("\t%v\n", v)
		}
	}
}
