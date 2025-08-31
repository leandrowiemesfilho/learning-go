package main

import "fmt"

/*
Demonstrate how a closure works.
In other words: create a function that returns another function, where this other function uses a variable beyond its scope.
*/

func main() {
	sum := getFuncSum(1, 2, 3)

	fmt.Println(sum())
}

func getFuncSum(nums ...int) func() int {
	return func() int {
		result := 0
		for _, num := range nums {
			result += num
		}

		return result
	}
}
