package main

import (
	"errors"
	"fmt"
)

func main() {
	array := [3]int {}
	number := 5
	array[0] = 3
	array[number-3] = 5

	fmt.Println(array)

	messages := []string{"1", "2", "3"}

	printMessages(messages)
}


func printMessages(messages []string) error {
	if len(messages) == 0 {
		return errors.New("Empty")
	}

	fmt.Println(messages)
	return nil
}