package main

import "fmt"

/*
Displays: The Unicode code point of all uppercase letters of the alphabet, three times each.
For example:
65
	U+0041 'A'
	U+0041 'A'
	U+0041 'A'
66
	U+0042 'B'
	U+0042 'B'
	U+0042 'B'
...and so on.
*/

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
