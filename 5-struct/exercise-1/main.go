package main

import "fmt"

/*
Create a "person" type with an underlying struct type that can contain the following fields:
First Name
Last Name
Favorite Ice Cream Flavors

Create two values of the "person" type and display these values,
using the range in the slice that contains the ice cream flavors.
*/

type person struct {
	firstName         string
	lastName          string
	favoriteIceCreams []string
}

func main() {
	p1 := person{
		firstName: "Derrick",
		lastName:  "Rose",
		favoriteIceCreams: []string{
			"Chocolate",
			"Gummy",
		},
	}

	p2 := person{
		firstName: "Michael",
		lastName:  "Jordan",
		favoriteIceCreams: []string{
			"Vanilla",
		},
	}

	showPersonDetails(p1)
	showPersonDetails(p2)
}

func showPersonDetails(p person) {
	fmt.Println(p.firstName, p.lastName)
	for _, v := range p.favoriteIceCreams {
		fmt.Printf("\t%v\n", v)
	}
}
