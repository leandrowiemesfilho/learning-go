package main

import "fmt"

/*
Create and use an anonymous struct.
Challenge: within the struct, have a value of type map and another of type slice.
*/

func main() {
	s := struct {
		m   map[int]string
		slc []string
	}{
		m: map[int]string{
			1: "a",
			2: "b",
		},
		slc: []string{
			"a",
			"b",
			"c",
		},
	}

	fmt.Println(s)
}
