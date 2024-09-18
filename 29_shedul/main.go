package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(0) // Limit to 2 OS threads

	for i := 0; i < 400000; i++ {
		go func(i int) {
			for j := 0; j < 5; j++ {
				fmt.Printf("Goroutine %d: %d\n", i, j)
				time.Sleep(100 * time.Millisecond) // Simulate work
			}
		}(i)
	}

	time.Sleep(1 * time.Second) // Allow time for goroutines to complete
}
