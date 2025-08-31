package main

import "fmt"

/*
Create a program that uses the switch statement,
where the switch statement is a string variable with identifier "favoriteSport".
*/

func main() {
	favoriteSport := "basket"

	switch favoriteSport {
	case "basket":
		fmt.Println("Lets try some jump shots")
	default:
		fmt.Println("Basket would be better")
	}
}
