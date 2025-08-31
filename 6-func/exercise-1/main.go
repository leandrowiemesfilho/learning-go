package main

import "fmt"

/*
Create a function that returns an int
Create another function that returns an int and a string
Call both functions
Demonstrate their results
*/

func main() {
	age := getAge()
	height, name := getHeightAndName()

	fmt.Println(height, name, age)
}

func getAge() int {
	return 32
}

func getHeightAndName() (int, string) {
	return 193, "Rose"
}
