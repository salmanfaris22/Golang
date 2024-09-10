package main

import (
	"fmt"
	"maps"
)

func main() {
	// m := make(map[string]int)
	// m["name"] = 12
	// m["age"] = 32
	// m["ages"] = 0
	// delete(m, "ages")
	// fmt.Println(m)
	// fmt.Println(len(m))

	m1 := map[string]int{"name": 30, "age": 32, "phone": 32}
	m2 := map[string]int{"name": 30, "age": 32, "phone": 32}

	fmt.Println(maps.Equal(m1, m2))

}
