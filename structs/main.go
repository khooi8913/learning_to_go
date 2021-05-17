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
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "jim@jim.com",
			zipCode: 11022,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("GG")

	// shortcut in Go
	jim.updateName("GG")

	jim.print()
}

// &variable - the memory address of the value that the variable is pointing to
// *pointer - the value in which the memory address is pointing at

// REFERENCE TYPES --- i.e., no need to worry about pointers
// slices, maps, channels, pointers, functions

// VALUE TYPES
// int, float, string, bool, structs

// Go ==> PASS BY VALUE

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
