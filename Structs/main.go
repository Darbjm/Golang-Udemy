package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// can use person{"Alex", "Anderson"} to assign structs but its a bad idea
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
