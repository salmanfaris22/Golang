package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter Rage Of number: ")
	fmt.Scan(&num)
	sum := 0
	for n := range num + 1 {
		if n%2 != 0 {
			sum += n
		}
	}
	fmt.Println(sum)

}
