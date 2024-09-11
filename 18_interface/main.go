package main

import "fmt"

type payment interface {
	Pay(n float32)
}

type gpay struct {
}

func (g gpay) Pay(a float32) {
	fmt.Println("Get Int G Pay :", a)
}

type paytm struct {
}

func (p paytm) Pay(a float32) {
	fmt.Println("Get Int Paytm :", a)
}

func main() {
	var NewPayment payment
	NewPayment = paytm{}
	NewPayment.Pay(32.0)

}

// import "fmt"

// type runner interface {
// 	walk()
// }

// type dog struct {
// 	name string
// }

// func (d dog) walk() {
// 	fmt.Println(d.name, "Is Runnig")
// }

// func main() {

// 	var run runner
// 	run = dog{name: "dog"}
// 	run.walk()

// }

// func main() {
// 	newPayment := payment{}
// 	newPayment.getPayment(1000)
// }

// type payment struct{}

// func (p payment) getPayment(amonut float32) {
// 	razerPaymentGEt := razerPay{}
// 	razerPaymentGEt.dealPay(amonut)
// }

// type razerPay struct{}

// func (r razerPay) dealPay(amonut float32) {
// 	fmt.Println("Done Payment Rs: ", amonut)
// }
