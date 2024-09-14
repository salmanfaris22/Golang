package main

import "fmt"

func somthing(n chan int) {

	t := 3
	x := 43
	n <- t
	n <- x
}

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}

	// done := make(chan int, 2)

	// go func() {
	// 	done <- 2
	// 	done <- 3
	// 	done <- 4
	// 	done <- 5
	// }()

	// get := <-done
	// get2 := <-done
	// get3 := <-done
	// get4 := <-done

	// fmt.Println(get, get2, get3, get4)

	// time.Sleep(time.Millisecond * 1000)

}
