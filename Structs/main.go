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

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
