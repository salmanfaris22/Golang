package main

import (
	"fmt"
)

func main() {
	var size string
	fmt.Println("enter string :")
	fmt.Scan(&size)
	fmt.Println(size)
	k := len(size) - 1

	for i := 0; i < len(size); i++ {
		if size[i] != size[k] {
			fmt.Print("not a palidrom")
			return
		}
		k--
	}
	fmt.Print("is a palidrom")
}
