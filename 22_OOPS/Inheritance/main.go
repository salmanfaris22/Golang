package main

import "fmt"

type Animal struct {
}

func (A Animal) Speek() string {

	return "Animal Is Speeking"
}

func main() {

	mic := Animal{}

	x := mic.Speek()
	fmt.Println(x)
}
