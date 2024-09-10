package main

import "fmt"

func counter() func() int {
	cont := 0

	return func() int {
		cont += 1
		return cont
	}
}

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

}
