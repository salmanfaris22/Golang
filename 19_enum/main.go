package main

import "fmt"

type status string

const (
	shipping status = "shipping"
	delivery        = "deliverd"
	exhanch         = "exchange"
	cancel          = "cancle"
)

func main() {
	payment(cancel)
}

func payment(a status) {
	fmt.Println("The Iteams Is : ", a)
}
