package main

import "fmt"

/*
Create a function that returns a function.
Assign the returned function to a variable.
Call the returned function.
*/

func main() {
	f := getFunc()

	f()
}

func getFunc() func() {
	return func() {
		fmt.Println("Hello World")
	}
}
