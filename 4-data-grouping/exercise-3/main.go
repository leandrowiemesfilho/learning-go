package main

import "fmt"

/*
Building on the previous exercise, use slicing to demonstrate the values:
From the first to the third item in the slice (including the third item!)
From the fifth to the last item in the slice (including the last item!)
From the second to the seventh item in the slice (including the seventh item!)
From the third to the second-to-last item in the slice (including the second-to-last item!)
Challenge: Obtain the same result as above using the len() function to determine the second-to-last item
*/

func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slc[:3])
	fmt.Println(slc[4:])
	fmt.Println(slc[1:7])
	fmt.Println(slc[2:9])
	fmt.Println(slc[2 : len(slc)-1])
}
