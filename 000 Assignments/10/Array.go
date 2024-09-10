package main

import "fmt"

func main() {

	var size int
	fmt.Println("Enter array Size")
	fmt.Scan(&size)
	arr1 := make([]int, size)
	arr2 := make([]int, size)

	fmt.Println("Enter Element Of Fisst array")
	for i := range size {
		fmt.Scan(&arr1[i])
	}
	fmt.Println("Enter Element Of Second array")
	for i := range len(arr2) {
		fmt.Scan(&arr2[i])
	}

	fmt.Println("array 1 is :", arr1)
	fmt.Println("array 1 is :", arr2)

}
