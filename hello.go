package main

import "fmt"

func main() {

	fmt.Println("Hello Wolrd")

	for i := 1; i <= 25; i++ {
		fmt.Print(i)
		if i%5 == 0 {
			fmt.Println()
		}
	}

}
