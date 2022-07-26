package main

import (
	"fmt"
	"khas/shape"
)

func main() {

	circle := shape.Circle{Radius: 5}
	fmt.Println(circle.Area())

	square := shape.NewSquare(5)
	fmt.Println(square.Area())
}