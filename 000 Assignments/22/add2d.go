package main

import "fmt"

func main() {
	var n int

	fmt.Println("Enter The Array length")
	fmt.Scan(&n)
	arr := make([][]int, n)
	setArry(&arr)

	AddArray(&arr)
	displayArray(arr)

}

func displayArray(arr [][]int) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func AddArray(arr *[][]int) {

	fmt.Println("Add Array element")
	for i := range *arr {
		for j := range (*arr)[i] {
			var n int
			fmt.Scan(&n)
			(*arr)[i][j] += n
		}
	}
}

func setArry(arr *[][]int) {
	fmt.Println("enter Array Element")

	for i := range *arr {
		(*arr)[i] = make([]int, len(*arr))
		for j := range (*arr)[i] {
			fmt.Scan(&(*arr)[i][j])
		}
	}

}
