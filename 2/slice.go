package main

import (
	"fmt"
)

func main() {
	var slice []int
	slices := make([]string, 3)
	some := append(slices, "hello", "Go")

	fmt.Println(some)

	fmt.Println(slice)
	fmt.Println(slices)
}