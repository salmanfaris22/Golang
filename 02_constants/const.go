package main

import "fmt"

func main() {
	const (
		name  string = "salman"
		lname string = "faris"
		age   int    = 3232
	)

	const num uint8 = 1
	// name = "salman"
	fmt.Println(name+lname, num)
}
