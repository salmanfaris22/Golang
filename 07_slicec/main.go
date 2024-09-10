package main

import (
	"fmt"
	"slices"
)

func main() {



	var num1 = []int{1, 2}
	var num2 = []int{1, 2}
	fmt.Println(slices.Equal(num1, num2))

}
