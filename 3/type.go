package main

import "fmt"

func main() {
	str := "hello"

	toStr(str)
	toStr(5)
}

func toStr(i interface{}) {
	// если присваивать без второго аргумента, то вызывается паника при ok = false
	str, ok := i.(string)
	fmt.Println(str)
	fmt.Println(ok)


}

// композиция интерфейсов

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}