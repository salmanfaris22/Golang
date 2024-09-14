package main

import "fmt"

type Animal struct {
	name string
}

func (A Animal) Speek() string {

	return "Animal Is Speeking"
}

type Dog struct {
	Animal
	breed string
}

func (D Dog) Speek() string {

	return "Dog Is Speeking"
}

func main() {

	mic := Dog{
		Animal: Animal{name: "cat"},
		breed:  "miciy",
	}
	fmt.Println(mic.name, mic.breed)
	x := mic.Speek()
	fmt.Println(x)
}
