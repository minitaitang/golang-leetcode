package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	m := make(map[string]bool)
	m["v-d/3_/v_4.mp4"] = true
	fmt.Print(m["v-d/3_/v_4.mp4"])
}
