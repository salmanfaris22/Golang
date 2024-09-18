package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("Hello")
	buf.WriteString("Hello")
	fmt.Println(buf.String())
}
