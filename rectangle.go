package main

import (
	"fmt"
)

type Rectangle struct {
	length, width float64
}

func NewRectangle(length, width float64) *Rectangle {
	return &Rectangle{length, width}
}

func (r *Rectangle) GetArea() float64 {
	return r.length * r.width
}

func (r *Rectangle) GetPerimeter() float64 {
	return (r.length + r.width) * 2
}

func main() {
	rect := NewRectangle(8, 5)
	fmt.Printf("Area of rectangle: %.2f\n", rect.GetArea())
	fmt.Printf("Perimeter of rectangle: %.2f\n", rect.GetPerimeter())
}
