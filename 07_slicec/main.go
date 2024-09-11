package main

import (
	"fmt"
	"slices"
)

func main() {

	var num1 = []int{1, 2}
	var num2 = []int{1, 2}
	arr := [2]int{2, 3}

	x := []int{2, 3}
	x = append(x, 3)
	fmt.Println(slices.Equal(num1, num2))
	fmt.Println(x, arr)

}
