package main

import "fmt"

type Contact struct {
	Email string
	Phone string
}

type Person struct {
	Firstname string
	Lastname  string
	Age       int
	Contact
}

func main() {
	p1 := Person{Firstname: "John", Lastname: "Da Costa", Age: 32}
	p2 := p1

	fmt.Printf("p1 == p2 -> %v\n", p1 == p2)

	p2.Age++

	fmt.Printf("p1 == p2 -> %v\n", p1 == p2)

	p3 := Person{
		Firstname: "John",
		Lastname:  "Da Costa",
		Age:       33,
		Contact: Contact{
			Email: "me@jdc.ch",
			Phone: "000 000 00 00",
		},
	}

	fmt.Printf("p2 == p3 -> %v\n", p2 == p3)

	p2.Contact = Contact{
		Email: "me@jdc.ch",
		Phone: "000 000 00 00",
	}

	fmt.Printf("p2 == p3 -> %v\n", p2 == p3)
}
