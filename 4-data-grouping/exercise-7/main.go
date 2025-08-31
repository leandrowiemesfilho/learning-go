package main

import "fmt"

/*
Create a slice containing string slices ([][]string). Assign values to this multidimensional slice as follows:
"First Name", "Last Name", "Favorite Hobby"
Include data for three people, and use range to display this data.
*/

func main() {
	multiSlice := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Dwayne", "Wade", "Play basketball"},
		{"Mister", "M", "Reveal magic tricks"},
	}

	for _, v := range multiSlice {
		for _, w := range v {
			fmt.Println(w)
		}
		fmt.Println()
	}
}
