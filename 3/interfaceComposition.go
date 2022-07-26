package main

import (
	"fmt"
	"math"
)


func main() {
	circle := Circle{8}
	square := Square{8}

	printArea(circle)
	printArea(square)

	printPerimeter(square)
	printAreaAndPerimeter(square)

}


// композиция интерфейсов

type ShapeWithAreaAndPerimeter interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func printPerimeter(p ShapeWithPerimeter) {
	fmt.Println(p.Perimeter())
}

func printAreaAndPerimeter(p ShapeWithAreaAndPerimeter) {
	fmt.Println(p.Area(), p.Perimeter())
}

func printArea(p ShapeWithArea) {
	fmt.Println(p.Area())
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

func (s Square) Perimeter() float32 {
	return s.sideLength * 4
}


type Circle struct {
	radius float32
}

func (s Circle) Area() float32 {
	return s.radius * s.radius * math.Pi
}
