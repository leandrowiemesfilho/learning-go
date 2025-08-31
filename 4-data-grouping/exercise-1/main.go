package main

import "fmt"

/*
Using a compound literal:
Create an array that supports 5 int values.
Assign values to its indexes.
Use range and display the array values.
Using format printing, display the array type.
*/

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i := range arr {
		fmt.Println(arr[i])
	}

	fmt.Printf("%T", arr)
}
