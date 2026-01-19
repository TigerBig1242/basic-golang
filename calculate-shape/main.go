package main

import (
	"fmt"
)

type Shapper interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (area Rectangle) Area() float64 {
	rectangleArea := area.Width * area.Height
	return rectangleArea
}

func (area Circle) Area() float64 {
	circleArea := area.Radius * area.Radius * (3.14)
	return circleArea
}

func PrintArea(shape Shapper) {
	fmt.Printf("Shape Area %f\n", shape.Area())
}

func main() {
	rectangle := Rectangle{7.42, 46.12}
	circle := Circle{6}

	PrintArea(rectangle)
	PrintArea(circle)
}
