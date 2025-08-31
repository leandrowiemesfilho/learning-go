package main

import "fmt"

/*
Create a program that uses the switch statement, without switch statement specified.
*/

func main() {
	result := 42

	switch {
	case result > 100:
		fmt.Println("result is greater than 100")
	case result == 42:
		fmt.Println("result is equal to 42")
	case result < 100:
		fmt.Println("result is less than 100")
	}
}
