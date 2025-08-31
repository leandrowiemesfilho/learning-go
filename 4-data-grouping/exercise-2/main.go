package main

import "fmt"

/*
Using a compound literal:
Create a slice of type int
Assign 10 values to it
Use range to display all these values.
And use format printing to display its type.
*/

func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range slc {
		fmt.Println(v)
	}

	fmt.Printf("%T", slc)
}
