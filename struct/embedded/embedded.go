package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (u Person) WhoAmI() {
	fmt.Printf("I'm %s\n", u.Name)
}

type User struct {
	Person
	Name string
	Role string
}

func (u User) WhoAmI() {
	fmt.Printf("I'm %s and my Role is %s\n", u.Name, u.Role)
}

func main() {
	groot := Person{Name: "Groot", Age: 20}
	groot.WhoAmI()

	admin := User{
		Person: Person{Name: "Johnny", Age: 31},
		Name:   "JDC",
		Role:   "SuperAdmin",
	}

	fmt.Println("call WhoAmI method")
	admin.Person.WhoAmI()
	admin.WhoAmI()

	fmt.Println("call name from type")
	fmt.Printf("I'm %s, from admin.Name\n", admin.Name)
	fmt.Printf("I'm %s, from admin.Person.Name\n", admin.Person.Name)

	// Embedded are not Inheritance
	var jdcPerson Person
	// jdcPerson = admin // ERR: cannot use admin (variable of type User) as Person value in assignment
	jdcPerson = admin.Person
	jdcPerson.WhoAmI()
}
