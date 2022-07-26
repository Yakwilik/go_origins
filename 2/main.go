package main

import "fmt"

// определяет название текущего пакета, в котором находится данный файл.
var m = 3

func init() {
	m = 5

}
// init – зарегистрированная функция, выполняется перед инициализацией пакета и перед вызовом функции main

func main() {

	fmt.Println(m)
	// function := increment()
	// function()
	// function()

	// function()
	// fmt.Println(function())
}

func increment() func() int {
	count := 0
	return func() (int) {  // принимает эта функция ничего и возвращает инт
		count++
		return count
	}
}