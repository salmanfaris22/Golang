package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("enter Name")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)

}
