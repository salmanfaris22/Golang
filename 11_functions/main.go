package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {

	result := add(4, 5)
	fmt.Println(result)
	fmt.Println(lnag())
}

func lnag() (string, string, string) {
	return "sabuth", "golanf", "sasi"
}
