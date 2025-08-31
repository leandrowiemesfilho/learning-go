package main

import "fmt"

/*
Create a new type: vehicle
The underlying type must be a struct
It must contain the following fields: doors, color
Create two new types: pickup and sedan
The underlying types must be structs
Both must contain "vehicle" as an embedded struct
The pickup truck type must contain a bool field called "fourWheel"
The sedan type must contain a bool field called "luxury"
Using the structs vehicle, pickup truck, and sedan:
Using a composite literal, create a value of type pickup truck and assign values to its fields
Using a composite literal, create a value of type sedan and assign values to its fields
Demonstrate these values.
*/

type vehicle struct {
	doors int
	color string
}

type pickup struct {
	fourWheel bool
	vehicle   vehicle
}

type sedan struct {
	luxury  bool
	vehicle vehicle
}

func main() {
	p := pickup{
		fourWheel: true,
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
	}

	s := sedan{
		luxury: false,
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
	}

	fmt.Println(p)
	fmt.Println(s)
}
