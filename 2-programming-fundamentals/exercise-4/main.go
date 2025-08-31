package main

import "fmt"

/*
Create a program that:
Assigns an int value to a variable
Expresses this value in decimal, binary, and hexadecimal
Shifts the bits of this variable by 1 to the left, and assigns the result to another variable
Expresses this other variable in decimal, binary, and hexadecimal
*/

func main() {
	x := 42

	fmt.Printf("decimal: %d\nbinary: %b\nhexadecimal: %#x\n", x, x, x)

	y := x << 1

	fmt.Printf("hexadecimal: %d\nbinary: %b\nhexadecimal: %d", y, y, y)
}
