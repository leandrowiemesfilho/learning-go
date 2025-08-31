package main

import "fmt"

/*
Callback: pass a function as an argument to another function.
*/

func main() {

	callFunc(sayHi)
}

func sayHi() {
	fmt.Println("Hello World")
}

func callFunc(f func()) {
	f()
}
