package main

import (
	"fmt"
	"sync"
)

/*
In addition to the main goroutine, create two additional goroutines.
Each additional goroutine should perform a separate printout.
Use waitGroups to ensure your goroutines finish before the program ends.
*/

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("Hello first routine")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("Hello second routine")
	}()

	wg.Wait()
	fmt.Println("All routines have been executed")
}
