package main

import (
	"fmt"
)

type price struct {
	totle int
}

type order struct {
	id   int
	name string
	age  int
	price
}

// func (o *order) changeAge(a int) {
// 	o.age = a
// }

func NewOrder(id int, name string, age int, totle int) *order {

	myOrder := order{
		id:   id,
		name: name,
		age:  age,
	}
	return &myOrder
}

func main() {
	// n := 0
	// lang := struct {
	// 	name string
	// }{"salman"}
	// getOrder := NewOrder(1, "salman", 21, 32)
	// fmt.Scan(&n)
	// getOrder.changeAge(n)

	name := order{
		id:   2,
		name: "salman",
		age:  32,
		price: price{
			totle: 3223,
		},
	}

	fmt.Println("GetPerson", name)
}
