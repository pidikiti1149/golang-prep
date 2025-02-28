package main

import "fmt"

type person struct {
	name string
	age  int
}

// Method with pointer receiver
func (p *person) updateName(newName string, newAge int) {
	p.name = newName
	p.age = newAge
}

// Method with value receiver
func (p person) showName() {
	fmt.Println("Name:", p.name)
}

func main() {
	a := person{name: "a", age: 25}
	a.showName()

	// Calling pointer method with value
	a.updateName("b", 67)
	fmt.Println("After pointer method:", a.name, a.age)

	// Calling value method with pointer
	a.showName()
}
