package main

import "fmt"

var employee = map[string]int{"age": 32, "salary": 32333}

func main() {
	employee["day"] = 32
	fmt.Println(employee)
}
