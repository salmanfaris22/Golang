// // Online Go compiler to run Golang program online
// // Print "Try programiz.pro" message

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	fmt.Println(diffWaysToCompute("2-1-1"))
// }

// func diffWaysToCompute(s string) bool {

// 	fmt.Println(s[:1])
// 	_, err := strconv.Atoi(s)

// 	return err == nil

// }

// // package main

// // import "fmt"

// // // type Dog struct {
// // // 	name  string
// // // 	breed string
// // // }

// // // type Cat struct {
// // // 	name string
// // // 	age  int
// // // 	typ  string
// // // }

// // // func (c Cat) Say() string {
// // // 	return "Myavooo"
// // // }

// // // func (d Dog) Say() string {
// // // 	fmt.Println(d.name)
// // // 	return "Bow Bow .."
// // // }

// // // type Animal interface {
// // // 	Say() string
// // // }

// // func main() {

// // 	Prim("fff")

// // 	// lizCat := Cat{
// // 	// 	name: "liza",
// // 	// 	age:  32,
// // 	// 	typ:  "pertion",
// // 	// }

// // 	// goger := Dog{
// // 	// 	name:  "bobi",
// // 	// 	breed: "indian",
// // 	// }

// // 	// GetInfo(goger)
// // 	// GetInfo(lizCat)
// // }

// // func Prim[T any](a T) {
// // 	fmt.Println(a)
// // }

// // // func GetInfo(a Animal) {
// // // 	fmt.Println(" Is Sya", a.Say())
// // // }

package main

import (
	"fmt"
)

func reverse(x int) int {
	s := 1
	if x < 0 {
		s = -1
		x = -x
	}

	r := 0
	for x != 0 {
		r = r*10 + x%10
		x /= 10
	}

	r *= s

	// if r < -1<<31 || r > (1<<31)-1 {
	// 	return 0
	// }
	return r
}

func main() {
	x := 123300
	fmt.Println(reverse(x)) // Output: 321
}
