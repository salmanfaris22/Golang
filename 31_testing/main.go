package main

import "fmt"

// type Dog struct {
// 	name  string
// 	breed string
// }

// type Cat struct {
// 	name string
// 	age  int
// 	typ  string
// }

// func (c Cat) Say() string {
// 	return "Myavooo"
// }

// func (d Dog) Say() string {
// 	fmt.Println(d.name)
// 	return "Bow Bow .."
// }

// type Animal interface {
// 	Say() string
// }

func main() {

	Prim("fff")

	// lizCat := Cat{
	// 	name: "liza",
	// 	age:  32,
	// 	typ:  "pertion",
	// }

	// goger := Dog{
	// 	name:  "bobi",
	// 	breed: "indian",
	// }

	// GetInfo(goger)
	// GetInfo(lizCat)
}

func Prim[T any](a T) {
	fmt.Println(a)
}

// func GetInfo(a Animal) {
// 	fmt.Println(" Is Sya", a.Say())
// }
