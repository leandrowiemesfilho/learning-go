package main

import "fmt"

/*
Create a struct "person"
Create a function called changeMe that takes *person as a parameter. This function should change a value stored at the address *person.
Tip: The "correct" way to dereference a value in a struct would be (*value).field
But there's an exception in the documentation. Link: https://golang.org/ref/spec#Selectors
"As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
→ x.f is shorthand for (*x).f." ←
In other words, we can use either the shortcut p1.name or the traditional (*p1).name
Create a value of type person;
Use the changeMe function and demonstrate the result.
*/

type person struct {
	name string
	age  int
}

func main() {
	p := person{
		name: "Bob",
		age:  42,
	}

	fmt.Println(p)

	changeMe(&p)

	fmt.Println(p)
}

func changeMe(p *person) {
	p.age = 45
}
