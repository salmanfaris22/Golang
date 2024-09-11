package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("how many elemnt In array")
	fmt.Scan(&size)
	arr := make([]int, size)
	fmt.Println("Enter Element Broh...!")
	for i := range size {
		fmt.Scan(&arr[i])
	}
	Sort(arr)
	fmt.Println(arr)
}

func Sort(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

}
