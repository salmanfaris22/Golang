package main

import "fmt"

func main() {
	name := ";l,;,"
	display(name)
}

func display[T any](b T) {
	fmt.Println("typed is", b)
}


