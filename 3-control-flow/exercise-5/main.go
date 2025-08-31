package main

import "fmt"

/*
Show the remainder when divided by 4 for all numbers between 10 and 100.
*/

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i, i%4)
	}
}
