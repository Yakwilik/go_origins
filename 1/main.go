package main

// определяет название текущего пакета, в котором находится данный файл.

import (
	"fmt"
	"reflect"
)

func main() {
	// print()
	// printHi()
	// printMessage("I am message")
	// printType("message")
	// pack.printHello()
	// a, b, c := 1, 2, "hello"
	// a, b, _ = b, a, c
	// fmt.Println(a, b, c)



	// message, check := enterTheClub(2)

	// if check != true {
	// 	fmt.Println("меня не впустили в клуб и сказали: \"" + message + "\"")
	// } else {
	// 	fmt.Println("меня впустили в клуб и сказали: \"" + message + "\"")
	// }
	fmt.Println(findMin(7, 5))

}


func print() {
	fmt.Printf("\nHello World %s", "Hello World\n")
}

func printHi() {
	message := "Hi"

	fmt.Println(message)

}

func printMessage(message string) {
	fmt.Println(message)
}

func printType(message string) {
	fmt.Println(reflect.TypeOf(message))
}

func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		return "входи, тебе можно", true
	}
	return "тебе еще нет 18", false
}

func findMin(numbers ... int) (int) {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if min > i {
			min = i
		}
	}

	return min
}