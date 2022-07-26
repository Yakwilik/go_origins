package main

import (
	"fmt"
)

func main()  {
	var m string = "hello"

	changeMessage(&m)
	fmt.Printf("m: %v\n", m)

}

func changeMessage(message *string) {
	*message += *message
}