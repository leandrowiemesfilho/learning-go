package main

import "fmt"

/*
Using the previous solution, place the values of type "person" in a map, using the last names as the key.
Demonstrate the map values using a range.
The different flavors should be demonstrated using another range, within the previous range.
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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for key, value := range m {
		fmt.Println(key)
		for _, v := range value.favoriteIceCreams {
			fmt.Printf("\t%v\n", v)
		}
	}
}
