package main

func main() {
	// defer откладывает выполнение функции в самый конец
	defer println("defer1")
	defer println("defer2")

	println("main()1")
	println("main()2")
	println("main()3")
}