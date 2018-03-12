package main

import (
	"fmt"
	"math"
)

// Shape interface. Abstract type
type Shape interface {
	// Here we are declaring behavior
	area() float64
}

// Square struct. Concret type
type Square struct {
	side float64
}

// Circle struct. Concret type
type Circle struct {
	radius float64
}

// Value receiver. Square implements Shape area behavior
func (s Square) area() float64 {
	return s.side * s.side
}

// Pointer receiver. Circle implements Shape area behavior
func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(&c)
}
