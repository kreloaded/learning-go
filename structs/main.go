package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Party",
		contact: contactInfo{
			email:   "alex@mymail.com",
			zipCode: 94000,
		},
	}

	// Pass by value
	// alex.print()
	// jim := alex.updateName("Jim")
	// jim.print()

	// Pass by reference
	alex.print()
	alexPointer := &alex
	alexPointer.updateName("Jim")
	alex.print()
}

func (personPointerToValue *person) updateName(newName string) {
	(*personPointerToValue).firstName = newName
}

func (p person) print() {
	fmt.Println(p)
	// fmt.Printf("%+v", p)
}
