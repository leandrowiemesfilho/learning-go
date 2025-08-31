package main

import "fmt"

/*
Assign a function to a variable.
Call that function.
*/

func main() {
	f := func() {
		fmt.Println("Hello World")
	}

	f()
}
