package main

import (
	"fmt"
	"math"
)

/*
Create a "square" type
Create a "circle" type
Create an "area" method for each type that calculates and returns the area of the figure
Circle area: 2 * π * radius
Square area: side * side
Create a "figure" type that defines as its interface any type that has the "area" method
Create an "info" function that takes a "figure" type and returns the area of the figure
Create a value of type "square"
Create a value of type "circle"
Use the "info" function to demonstrate the area of the "square"
Use the "info" function to demonstrate the area of the "circle"
*/

type figure interface {
	area() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	s := square{side: 5}
	c := circle{radius: 5}

	info(s)
	info(c)
}

func info(f figure) {
	fmt.Println(f.area())
}
