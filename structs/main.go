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

	alex.print()
	jim := alex.updateName("Jim")
	jim.print()
}

func (p person) updateName(newName string) person {
	p.firstName = newName

	return p
}

func (p person) print() {
	fmt.Println(p)
	// fmt.Printf("%+v", p)
}
