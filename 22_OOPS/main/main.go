package main

import (
	"fmt"

	"main/struct/person.go" // replace with your module path
)

func main() {
	p := person.NewPerson("John Doe", 30)
	fmt.Println("Name:", p.GetName())
	fmt.Println("Age:", p.GetAge())

	p.SetName("Jane Doe")
	p.SetAge(25)

	fmt.Println("Updated Name:", p.GetName())
	fmt.Println("Updated Age:", p.GetAge())
}
