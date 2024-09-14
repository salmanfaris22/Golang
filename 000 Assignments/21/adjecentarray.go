package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter The Array length")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter The Array elemnt")
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	newArr := adjecent(arr)
	fmt.Println(newArr)
}

func adjecent(arr []int) []int {
	n := make([]int, len(arr)-1)
	for i := range n {
		n[i] = arr[i] * arr[i+1]
	}
	return n
}
