package main

import "fmt"

/*
Create a type for a struct called "person"
Create a "talk" method for this type that has a receiver pointer (*person)
Create an interface, "humans," that is implemented by types with the "talk" method
Create a "saySomething" function whose parameter is of type "humans" and that calls the "talk" method
Demonstrate in your code:
That you can use a value of type "*person" in the "saySomething" function
That you cannot use a value of type "person" in the "saySomething" function
*/

type person struct {
	name string
}

func (p *person) talk() {
	fmt.Println(p.name)
}

type humans interface {
	talk()
}

func saySomething(h humans) {
	h.talk()
}

func main() {
	p := &person{
		name: "John Doe",
	}

	saySomething(p)

	/*
		This block will fail, because the pointer to person is implemented by the interface humans,
		not the person type itself
		p1 := person{
			name: "John Doe",
		}

		saySomething(p1)
	*/
}
