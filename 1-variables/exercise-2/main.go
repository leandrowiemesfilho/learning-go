package main

import "fmt"

/*
Use var to declare three variables. They must have package-level scope. Do not assign values to these variables.
Use the following identifiers and types for these variables:
	1. Identifier "x" must have type int
	2. Identifier "y" must have type string
	3. Identifier "z" must have type bool
In the main function:
	1. Show the values of each identifier
	2. The compiler assigned values to these variables. What are these values called?
*/

var x int
var y string
var z bool

func main() {
	fmt.Println("x variable:", x)
	fmt.Println("y variable:", y)
	fmt.Println("z variable:", z)
}
