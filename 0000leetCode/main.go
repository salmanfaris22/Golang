package main

import "fmt"

func main() {
	t := 7

	arr := []int{1, 3, 5, 6}
	x := searchInsert(arr, t)
	fmt.Println(x)
}

func searchInsert(nums []int, target int) int {
	s := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < target {

			s = i

		}
	}

	return s + 1
}
