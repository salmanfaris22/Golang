package main

import "fmt"

func main() {
	name := 32
	display(name)
}

func display[T any](b T) {
	fmt.Println("typed is", b)
}
