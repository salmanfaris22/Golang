package main

import "fmt"

type Array struct {
	array [][]int
	row   int
	col   int
}

func main() {
	var c, r int
	fmt.Println("Enter The Array Row")
	fmt.Scan(&r)
	fmt.Println("Enter The Array Column ")
	fmt.Scan(&c)

	arr := Array{
		array: make([][]int, c),
		row:   r,
		col:   c,
	}
	for i := range arr.array {
		arr.array[i] = make([]int, r)
	}
	fmt.Println("Enter Aray element")
	arr.setArray()
	arr.display()

}

func (a Array) setArray() {
	for i := range a.array {
		for j := range a.array[i] {
			fmt.Scan(&a.array[i][j])
		}
	}
}
func (a Array) display() {
	for i := range a.array {

		fmt.Println(a.array[i])

	}
}
