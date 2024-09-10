package main

import (
	"fmt"
)

func GetNumber() float32 {
	var n1, n2 float32
	fmt.Println("enter Num 1")

	_, err := fmt.Scan(&n1)
	if err != nil {
		fmt.Println("Pleas Enter A valide Number")
		fmt.Scanln()
		return GetNumber()
	}
	fmt.Println("enter Num 2")
	_, err2 := fmt.Scan(&n2)

	if err2 != nil {
		fmt.Println("Pleas Enter A valide Number")
		fmt.Scanln()
		return GetNumber()
	}

	return n1 + n2
}

func main() {

	x := GetNumber()
	fmt.Println("The Sum Is", x)

}
