package main

import "fmt"

/*
Using the solution from the previous exercise:
1. In package-level scope, using the var keyword, create a variable with the identifier "y." The type of this variable must be the underlying type of the type you created in the previous exercise.
2. In the main function:
	1. This should already be done:
		1. Prove the value of the variable "x"
		2. Prove the type of the variable "x"
		3. Assign 42 to the variable "x" using the "=" operator
		4. Prove the value of the variable "x"
	2. Now also do:
		1. Use conversion to transform the type of the value of the variable "x" into its underlying type and, using the "=" operator, assign the value of "x" to "y"
		2. Prove the value of "y"
		3. Prove the type of "y"
*/

type myInt int

var x myInt
var y int

func main() {
	fmt.Printf("x value: %v\nx type: %T\n", x, x)

	x = 42

	fmt.Printf("new x value %v\n", x)

	y = int(x)

	fmt.Printf("y value: %v\ny type: %T", y, y)
}
