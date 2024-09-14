package main

import "fmt"

type Maths struct{}

func (M Maths) addition(n1, n2 float32) float32 {
	return n1 + n2
}
func (M Maths) subtraction(n1, n2 float32) float32 {
	return n1 - n2
}

func (M Maths) multiplication(n1, n2 float32) float32 {
	return n1 * n2
}

func (M Maths) division(n1, n2 float32) float32 {
	return n1 - n2
}

func main() {
	var s, n1, n2, result float32
	var calc Maths
	fmt.Println("Select The Typ :")
	fmt.Println("1: addition")
	fmt.Println("2: subtraction:")
	fmt.Println("3: multiplication :")
	fmt.Println("4: division:")
	fmt.Scan(&s)
	if s < 0 || 4 <= s {
		fmt.Println("Invalid choice. Please try again.")

	}
	fmt.Println("Enter Your firsrt Number")
	fmt.Scan(&n1)
	fmt.Println("Enter Your Second Number")
	fmt.Scan(&n2)

	switch {
	case s == 1:
		result = calc.addition(n1, n2)
	case s == 2:
		result = calc.subtraction(n1, n2)
	case s == 3:
		result = calc.multiplication(n1, n2)
	case s == 4:
		result = calc.division(n1, n2)

	default:
		fmt.Println("enter envelid Number")
	}
	fmt.Println("the Result Is :", result)
}
