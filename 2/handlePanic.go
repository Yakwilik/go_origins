package main

import "fmt"


func main() {

	defer handlePanic()
	messages := []string{
		"0",
		"1",
		"2",
	}
	fmt.Println(messages)

	messages[3] = "3"

	fmt.Println(messages)
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}