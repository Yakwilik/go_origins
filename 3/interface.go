package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (s Circle) Area() float32 {
	return s.radius * s.radius * math.Pi
}

func main() {
	circle := Circle{8}
	square := Square{8}

	printArea(circle)
	printArea(square)
	printInterface(circle)
	printInterface(square)

	println("\n\n\n")

	printInterfaceType(5)
	printInterfaceType("fasd")
	printInterfaceType(true)
	printInterfaceType(circle)
	printInterfaceType(square)
}

// интерфейсная функция печати площади фигуры
func printArea(shape Shape) {
	fmt.Println(shape.Area())
}

func printInterface(i interface{}) {
	fmt.Printf("%+v \n", i)
}

func printInterfaceType(i interface{}) {
	switch value := i.(type) {
	case int:
		println("int", value)
	case bool:
		println("boolean", value)
	case string:
		println("string", value)
	default:
		println("unknown type", value)
	}

}