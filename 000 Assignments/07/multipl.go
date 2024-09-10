package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter Number For MultiPlication : ")
	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%dX%d=%d\n", i, num, i*num)
	}
}
