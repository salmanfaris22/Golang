package main

import (
	"fmt"
)

type order struct {
	id   int
	name string
	age  int
}

func (o *order) changeAge(a int) {
	o.age = a
}

func NewOrder(id int, name string, age int) *order {

	myOrder := order{
		id:   id,
		name: name,
		age:  age,
	}
	return &myOrder
}

func main() {
	n := 0
	lang := struct {
		name string
	}{"salman"}
	getOrder := NewOrder(1, "salman", 21)
	fmt.Scan(&n)
	getOrder.changeAge(n)

	fmt.Println("GetPerson", getOrder, lang)
}
