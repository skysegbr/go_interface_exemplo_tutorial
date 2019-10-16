package main

import (
	"math"
)

// Circle - "class"
type Circle struct {
	Radius float64
	Qux string // Obriga o uso explicitamente "nomeado" {Radius: 5} não permite {5}
}

// Area of circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perim of circle
func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}