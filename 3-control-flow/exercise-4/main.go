package main

import "fmt"

/*
Create a loop using the syntax: for {}
Use it to show the years since you were born.
*/

func main() {
	i := 2000

	for {
		if i > 2025 {
			break
		}

		fmt.Println(i)
		i++
	}
}
