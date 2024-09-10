package main

import "fmt"

func main() {

	fmt.Println("Enter Your Mark Bro ...")

	result := Mark()

	fmt.Println(result)
}

func Mark() string {
	var num float32
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Pleas Enter A mark")
		fmt.Scanln()
		return Mark()
	}

	if num >= 101 || num < 0 {
		fmt.Println("Pleas Enter A valide Mark")
		fmt.Scanln()
		return Mark()
	}

	if num >= 50 {
		return "congrats your Pass The Exam"
	} else {
		return "Sry Bro your Loos The Exam"
	}

}
