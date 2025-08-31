package main

import "fmt"

/*
Create a value and assign it to a variable.
Show the address of this value in memory.
*/

func main() {
	i := 42

	fmt.Println(&i)
}
