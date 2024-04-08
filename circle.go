package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
	color  string
}

func NewCircle() *Circle {
	return &Circle{radius: 1.0, color: "red"}
}

func NewCircleWithRadius(r float64) *Circle {
	return &Circle{radius: r, color: "red"}
}

func (c *Circle) GetRadius() float64 {
	return c.radius
}

func (c *Circle) GetArea() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c1 := NewCircle()
	fmt.Printf("The circle has radius of %f and area of %f\n", c1.GetRadius(), c1.GetArea())

	c2 := NewCircleWithRadius(2.0)
	fmt.Printf("The circle has radius of %f and area of %f\n", c2.GetRadius(), c2.GetArea())
}
