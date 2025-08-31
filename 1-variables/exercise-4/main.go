package main

import "fmt"

/*
Create a type. The underlying type must be int.
Create a variable for this type, with the identifier "x," using the var keyword.
In the main function:
	1. Set the value of the variable "x"
	2. Set the type of the variable "x"
	3. Assign 42 to the variable "x" using the "=" operator
	4. Set the value of the variable "x"

To keep learning: https://go.dev/ref/spec#Types
*/

type myInt int

var x myInt

func main() {
	fmt.Printf("x value: %v\nx type: %T\n", x, x)

	x = 42

	fmt.Printf("new x value %v", x)
}
