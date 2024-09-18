package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	tow := make(chan string)
	quti := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		one <- "hy"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		tow <- "Hello"
	}()

	for {
		select {
		case rec1 := <-one:
			fmt.Println("start One:", rec1)
		case rec2 := <-tow:
			fmt.Println("Start Two:", rec2)
		case <-quti:
			return
		}
	}

}
