package main

import "fmt"

type contactInfo struct {
	email    string
	zipeCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

// When changing things in a function remember the difference bewteen value types and reference types.
// Check image for reference

func main() {
	// // can use person{"Alex", "Anderson"} to assign structs but its a bad idea
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// // to have intialize as empty use
	// var jack person
	// jack.firstName = "jack"
	// jack.lastName = "nickels"

	// fmt.Println(jack)
	// fmt.Printf("%+v", jack)

	jim := person{
		firstName: "Jim",
		lastName:  "Dun",
		contactInfo: contactInfo{
			email:    "jim@gmail.com",
			zipeCode: 9400,
		},
	}

	// points to ram
	jimPointer := &jim // could just use the line below as a shortcut
	jimPointer.updateName("jimmy")
	jim.print()
}

// pointer to person
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
