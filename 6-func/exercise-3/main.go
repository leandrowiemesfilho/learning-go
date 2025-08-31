package main

import "fmt"

/*
Use the defer statement in a way that demonstrates
that its execution only occurs at the end of the context to which it belongs.
*/

func main() {
	defer fmt.Println("Last thing to be executed here")

	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}
