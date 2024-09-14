package main

import "fmt"

type income struct {
	income float32
}

func main() {
	var s float32
	fmt.Println("Ente Your Anuval Salary")
	fmt.Scan(&s)
	salary := income{
		income: s,
	}
	salary.Tax()

}

func (i income) Tax() {
	var tax float32

	switch {
	case i.income < 250000:

		fmt.Println("No Tax avalible", tax)

	case i.income <= 500000:

		tax = i.income * 5 / 100
		fmt.Println(" tax avalible is:", tax)

	case i.income <= 1000000:
		tax = i.income * 20 / 100
		fmt.Println(" Tax avalible is:")

	case i.income <= 5000000:
		tax = i.income * 30 / 100
		fmt.Println(" Tax avalible is:")

	}
}
