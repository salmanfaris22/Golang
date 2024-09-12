package main

import "fmt"

func makeArray(b int) [][]int {

	arr := make([][]int, b)

	for i := 0; i < b; i++ {
		arr[i] = make([]int, b)
		for j := 0; j < b; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	return arr
}

func sumArray(arr1, arr2 [][]int) [][]int {
	arr := make([][]int, len(arr1))
	for i := 0; i < len(arr2); i++ {
		arr[i] = make([]int, len(arr))
		for j := 0; j < len(arr1); j++ {
			arr[i][j] = arr1[i][j] + arr2[i][j]
		}
	}
	return arr
}

func main() {

	size := 0
	fmt.Printf("Enter Your Array Size")

	fmt.Scan(&size)
	fmt.Printf("Enter First array element")
	arr1 := makeArray(size)
	fmt.Println("Enter Seond Aarray Element")
	arr2 := makeArray(size)
	sum := sumArray(arr1, arr2)
	fmt.Println(sum)

}
