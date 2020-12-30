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

	// points to ram
	jimPointer := &jim
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
