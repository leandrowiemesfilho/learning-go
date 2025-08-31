package main

import "fmt"

/*
Create a string variable using a raw string literal.
Demonstrate it.
*/

func main() {
	x := `
			SELECT
				* 
			FROM user;
`

	fmt.Println(x)
}
