package main

import "fmt"

func main() {
	defer fmt.Println("one")
	defer fmt.Println("tow")
	defer fmt.Println("tree")

}
