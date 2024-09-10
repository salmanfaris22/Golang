package main

import "fmt"

func main() {
	fmt.Println("one")

	panic("pamoker")
	// defer recover()
	fmt.Println("tree")

}
