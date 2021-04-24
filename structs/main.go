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

}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v", p)
}
