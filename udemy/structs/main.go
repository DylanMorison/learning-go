package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@hjim.com",
			zipCode: 91400,
		},
	}

	// give me access to the address in memeory that jim is at 
	// jimPointer is a reference to the memory address that jim is at
	// jimPointer := &jim

	jim.updateName("New First Name")
	jim.print()
}

// go a pass by value language
// meaning whenever we pass a value into a function
// go takes that value and places it into a new container
// in memory.  So in this function we are actually creating a 
// new value of person in memory with a new firstName 

// go shortcut: we can call this method with either a pointer to person type
// or a value of person type
func (pointerToPerson *person) updateName(newFirstName string) {
	// a star in front of a type means something completely different then a star
	// in front of a person. 
	// When in front of a type it means were working with a poiner to that type
	// When in front of a value, it means we want to manipulate the value the pointer
	// is referencing.
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
