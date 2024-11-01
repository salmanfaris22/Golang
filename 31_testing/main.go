// package main

// import (
// 	"errors"
// 	"fmt"
// )

// type Hey struct {
// }

// func (H Hey) Say() { //recever metherd
// 	fmt.Println("hyyy")

// }

// func main() {
// 	h := Hey{}
// 	h.Say()

// 	x := 100

// 	PrintValue(&x)

// 	num, err := Modfy(nil)
// 	if err != nil {
// 		return
// 	}
	
// 	fmt.Println(x)
// }

// func PrintValue(a *int) {
// 	fmt.Println("function Hy")
// }

// func Modfy(x *int) (*int, error) {
// 	if x != nil {
// 		*x++
// 		return x, nil
// 	}
// 	return nil, errors.New("err")
// 	// defefer
// }

// //  stack trace
// // default value  types
// //recever metherd
// //










// package main

// import (
//     "fmt"
//     "time"
// )

// func task(id int) {
//     fmt.Printf("Task %d starting\n", id)
//     time.Sleep(time.Second) // Simulate work
//     fmt.Printf("Task %d done\n", id)
// }

// func main() {
//     for i := 1; i <= 3; i++ {
//         go task(i) // Start tasks concurrently
//     }
    
//     time.Sleep(2 * time.Second) // Wait for tasks to finish
//     fmt.Println("All tasks completed")
// }


package main

import (
    "fmt"
    "runtime"
    "sync"
)

func task(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Task %d starting\n", id)
    // Simulate heavy computation
    for i := 0; i < 100000000; i++ {}
    fmt.Printf("Task %d done\n", id)
}

func main() {
    runtime.GOMAXPROCS(4) // Allow up to 4 threads to run concurrently
    var wg sync.WaitGroup

    for i := 1; i <= 4; i++ {
        wg.Add(1)
        go task(i, &wg) // Start tasks in parallel
    }
    
    wg.Wait() // Wait for all tasks to finish
    fmt.Println("All tasks completed")
}
