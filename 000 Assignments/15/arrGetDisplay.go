package main

import "fmt"

func setArray(arr []int) {
	fmt.Println("Enter element")
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

}

func display(arr []int) {
	fmt.Println("The element Is : ")
	for i := range len(arr) {
		fmt.Print(arr[i], " ")
	}
}

func main() {
	size := 0
	fmt.Println("How Many Element Element")
	fmt.Scan(&size)
	arr := make([]int, size)

	setArray(arr)

	display(arr)

}
