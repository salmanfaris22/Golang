package main

import "fmt"

func main() {

	// var nums []int

	var nums = make([]int, 3, 5)
	// nums = append(nums, 3)
	// nums = append(nums, 3)
	// nums = append(nums, 3)
	// nums = append(nums, 3)

	// nums := []int{}
	nums[0] = 2
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

}
