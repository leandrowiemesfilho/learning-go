package main

import "fmt"

/*
Using iota, create four constants whose values are the next four years.
Demonstrate these values.
*/

const (
	_ = 2025 + iota
	a
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
