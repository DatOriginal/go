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

func NewCircleWithColor(color string) *Circle {
	return &Circle{radius: 1.0, color: color}
}

func NewCircleWithColorAndRadius(radius float64, color string) *Circle {
	return &Circle{radius: radius, color: color}
}

func (c *Circle) GetRadius() float64 {
	return c.radius
}

func (c *Circle) GetArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c * Circle) GetColor() string {
	return c.color
}

func main() {
	c1 := NewCircle()
	fmt.Printf("The circle has radius of %f and area of %f and color is %s\n", c1.GetRadius(), c1.GetArea(), c1.GetColor())

	c2 := NewCircleWithRadius(2.0)
	fmt.Printf("The circle has radius of %f and area of %f and color is %s\n", c2.GetRadius(), c2.GetArea(), c2.GetColor())

	c3 := NewCircleWithColor("black")
	fmt.Printf("The circle has radius of %f and area of %f and color is %s\n", c3.GetRadius(), c3.GetArea(), c3.GetColor())

	c4 := NewCircleWithColorAndRadius(2.0, "black")
	fmt.Printf("The circle has radius of %f and area of %f and color is %s\n", c4.GetRadius(), c4.GetArea(), c4.GetColor())
}
