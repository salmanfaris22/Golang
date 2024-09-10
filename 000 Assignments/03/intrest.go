package main

import "fmt"

func main() {

	var amount int
	var intre, year float32

	fmt.Println("Enter Your Principal amount :")
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Enter valid amount")

	}
	fmt.Println("Enter Your Interest rate :")
	_, err2 := fmt.Scan(&intre)
	if err2 != nil {
		fmt.Println("Enter valid intrest")

	}
	fmt.Println("Enter Your Number of years :")
	_, err3 := fmt.Scan(&year)
	if err3 != nil {
		fmt.Println("Enter valid year")

	}
	x := intrest(amount, intre, year)
	fmt.Println("your inters for", year, "is :", x)
}

func intrest(p int, r, n float32) float32 {
	newP := float32(p)
	return (newP * r * n) / 100
}
