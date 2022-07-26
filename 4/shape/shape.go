package shape

import "math"


type Shape interface {
	Area() float32
}

type Circle struct {
	Radius float32
}

func (a Circle) Area() float32 {
	return a.Radius * a.Radius * math.Pi
}

type Square struct {
	sideLength float32
}

func (a Square) Area() float32 {
	return a.sideLength * a.sideLength
}

func NewSquare(sideLength float32) (Square) {
	return Square{
		sideLength: sideLength,
	}
}