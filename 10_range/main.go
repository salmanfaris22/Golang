package main

import "fmt"

func main() {
	// nums := []int{6, 7, 8, 9}
	// m1 := map[string]int{"name": 30, "age": 32, "phone": 32}
	for i, num := range "slaman" {
		fmt.Println("value is", string(num), "key is", i)
	}
}
