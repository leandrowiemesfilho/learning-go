package main

import "fmt"

/*
Create a function that receives a variadic parameter of type int and returns the sum of all ints received;
Pass a slice of int value as an argument to the function;
Create another function; this one should receive a slice of int value and return the sum of all elements of the slice;
Pass a slice of int value as an argument to the function.
*/

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(nums...))
	fmt.Print(sumSlice(nums))
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}

func sumSlice(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}
