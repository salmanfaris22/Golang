package main

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

func getFuns(fn func(a, b int) int) (int, string) {

	return fn(4, 6), "salman"

}

func main() {
	fn := func(a, b int) int {
		return a + b
	}

	x, s := getFuns(fn)
	fmt.Println(x, s)

	// result := add(4, 5)
	// a, b, c := lnag()
	// fmt.Println(result)
	// fmt.Println(a, b, c)
}

// func lnag() (string, string, string) {
// 	return "sabuth", "golanf", "sasi"
// }
