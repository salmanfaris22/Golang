package main

import "fmt"

func main() {
	num := 5

	changeNum(&num)

	fmt.Println(num)

}

func changeNum(n *int) {
	*n = 19
	fmt.Println(*n)
}
