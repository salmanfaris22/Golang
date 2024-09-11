package main

import "fmt"

func main() {
	var size int
	fmt.Println("how many elemnt In array")
	fmt.Scan(&size)
	arr := make([]int, size)
	fmt.Println("Enter Element Broh...!")
	for i := range size {
		fmt.Scan(&arr[i])
	}
	e := Even(arr)
	fmt.Printf("Ther Have %d even number", e)
}

func Even(num []int) int {
	e := 0
	for i := range len(num) {
		if num[i]%2 == 0 {
			e++
		}
	}

	return e
}
