package main

import "fmt"

func printNumbers(n int) {
	fmt.Println(n)
}


func main() {
	for i := 0; i < 10; i++ {
		go printNumbers(i)
	}
}
