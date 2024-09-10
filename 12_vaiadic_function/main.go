package main

import "fmt"

func sum(num ...int) int {
	res := 0

	for _, i := range num {
		res += i
	}

	return res
}

func main() {

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	res := sum(nums...)
	fmt.Println(res)

}
