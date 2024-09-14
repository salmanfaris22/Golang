package main

import "fmt"

func main() {
	var w, l, a float32
	fmt.Println("Enter Marks")

	fmt.Println("Written test = ")
	fmt.Scan(&w)
	fmt.Println("Lab exams =  ")
	fmt.Scan(&l)
	fmt.Println("Assignments = ")
	fmt.Scan(&a)

	avarege := ((w * 70) / 100) + ((l * 20) / 100) + ((a * 10) / 100)
	fmt.Println(avarege)
}
