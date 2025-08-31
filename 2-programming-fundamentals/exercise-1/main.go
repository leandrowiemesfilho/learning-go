package main

import "fmt"

/*
Write a program that displays a number in decimal, binary, and hexadecimal.
*/

func main() {
	x := 42
	fmt.Printf("decimal: %d\nbinary: %b\nhex: %#x", x, x, x)
}
