package main

import "fmt"

/*
Using the shorthand operator, assign these values to variables with identifiers "x", "y", and "z".
 1. 42
 2. "James Bond"
 3. true

Now, demonstrate the values in these variables with:
 1. A single print statement
 2. Multiple print statements
*/

func main() {
	x := 42
	y := "James Bond"
	z := true

	singleStatement(x, y, z)
	multipleStatements(x, y, z)
}

func singleStatement(x int, y string, z bool) {
	fmt.Println(x, y, z)
}

func multipleStatements(x int, y string, z bool) {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
