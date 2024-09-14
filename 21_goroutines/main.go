package main

import (
	"fmt"
	"sync"
)

func tas(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(n)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go tas(i, &wg)
	}
	wg.Wait()
}
