package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	shakti := person{
		firstName: "Shakti",
		lastName:  "Ratan",
		contact: contactInfo{
			email:   "satrat@gmail.con",
			zipcode: 123321,
		}}
	shaktiPointer := &shakti
	shaktiPointer.updateName("sat")
	shakti.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
