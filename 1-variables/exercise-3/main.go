package main

import "fmt"

/*
Using the solution from the previous exercise:
1. At package-level scope, assign the following values to the variables:
	1. For "x", assign 42
	2. For "y", assign "James Bond"
	3. For "z", assign true
2. In the main function:
	1. Use fmt.Sprintf to assign all these values to a single variable. Make this string assignment to a variable named "s" using the shorthand operator.
	2. Demonstrate the variable "s".
*/

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("x:%d; y:%s, z:%v", x, y, z)

	fmt.Println(s)
}
