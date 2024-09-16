package main

import "fmt"

func main() {
	num := [4]int{1, 2}
	num[1] = 10
	num[3] = 23
	fmt.Println(num)
}

