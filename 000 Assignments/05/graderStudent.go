package main

import "fmt"

func main() {
	var mark int
	fmt.Println("Pleas Enter Your Mark")
	fmt.Scan(&mark)
	ChackGred(mark)
}

func ChackGred(mark int) {
	switch {
	case mark > 90:
		fmt.Println("Your A gred Student ")
	case mark > 80:
		fmt.Println("Your B gred Student ")
	case mark > 70:
		fmt.Println("Your C gred Student ")
	case mark > 60:
		fmt.Println("Your D gred Student ")
	case mark > 50:
		fmt.Println("Your E gred Student ")
	default:
		fmt.Println("Your faild Bro ")
	}
}
