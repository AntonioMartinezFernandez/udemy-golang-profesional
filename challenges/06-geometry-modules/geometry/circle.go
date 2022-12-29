package geometry

import (
	"math"
)

// Circle Class
type Circle struct {
	Radius float32
}

func (circle *Circle) calculateArea() float32 {
	return math.Pi * (circle.Radius * circle.Radius)
}

func (circle *Circle) calculatePerimeter() float32 {
	return 2 * math.Pi * circle.Radius
}
