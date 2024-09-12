package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter Your Number")
	fmt.Scan(&num)

	x := checkPrime(num)

	if x {
		fmt.Println("This Is A Prime Number")
	} else {
		fmt.Println("this is Not A prmine Nmnber")
	}
}

func checkPrime(n int) bool {
	if n == 2 || n == 1 {
		return true
	}
	for i := 3; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
