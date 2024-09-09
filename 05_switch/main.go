package main

import "fmt"

func main() {
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("default")
	// }
	// fmt.Println(time.Now().Weekday())
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("WeekEnd")
	// default:
	// 	fmt.Println("WeekDay")
	// }

	WhoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Print("its intiger")
		case bool:
			fmt.Println("its boolina")
		case string:
			fmt.Println("its String")
		default:
			fmt.Println(i)
		}
	}

	WhoAmI("as")
}
