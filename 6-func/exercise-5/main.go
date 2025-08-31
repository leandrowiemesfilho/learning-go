package main

import "fmt"

/*
Create and use an anonymous function.
*/

func main() {
	func() {
		fmt.Println("Hello World")
	}()
}
