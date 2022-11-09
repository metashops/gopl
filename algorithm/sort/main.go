package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["a"] = "A"
	m["b"] = "B"
	if _, ok := m["a"]; !ok {
		m["c"] = "C"
	}
	fmt.Println(m)

}
