package main

import (
	"fmt"
	"slices"
)

func main() {

	// var nums []int

	// var nums = make([]int, 3, 5)
	// nums = append(nums, 3)
	// nums = append(nums, 3)
	// nums = append(nums, 3)
	// nums = append(nums, 3)

	// nums := []int{}
	// nums[0] = 2
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// var num1 = make([]int, 0, 5)
	// num1 = append(num1, 2)
	// var num2 = make([]int, len(num1))
	// fmt.Println(num2, num1)
	// copy(num2, num1)
	// fmt.Println(num2, num1)

	// var nums = []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(nums[:4])

	var num1 = []int{1, 2}
	var num2 = []int{1, 2}
	fmt.Println(slices.Equal(num1, num2))

}
