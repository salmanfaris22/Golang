package main

import (
    "fmt"
    "sync"
    "time"
)

func printNumbers(wg *sync.WaitGroup, n int) {
    defer wg.Done() // Notify that this goroutine is done
    for i := 1; i <= n; i++ {
        fmt.Println(i)
        time.Sleep(100 * time.Millisecond) // Simulate work
    }
}

func main() {
    var wg sync.WaitGroup

    wg.Add(2) // Add two goroutines to wait for

    go printNumbers(&wg, 5) // Start first goroutine
    go printNumbers(&wg, 5) // Start second goroutine

    wg.Wait() // Wait for both goroutines to finish
    fmt.Println("All goroutines complete.")
}
