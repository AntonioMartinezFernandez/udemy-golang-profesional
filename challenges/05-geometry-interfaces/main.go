package main

import (
	"fmt"
	"math"
)

func main() {
	rectangle := Rectangle{
		width:  5.2,
		height: 10.6,
	}

	circle := Circle{
		radius: 8.1,
	}

	printGeometricMeasures(&rectangle)
	printGeometricMeasures(&circle)
}

/*
Interface
*/
type Geometry interface {
	calculateArea() float32
	calculatePerimeter() float32
}

func printGeometricMeasures(g Geometry) {
	fmt.Println("Area: ", g.calculateArea())
	fmt.Println("Perimeter: ", g.calculatePerimeter())
}

/*
	Structs (classes) and their methods
*/

// Rectangle
type Rectangle struct {
	width  float32
	height float32
}

func (rectangle *Rectangle) calculateArea() float32 {
	return rectangle.width * rectangle.height
}

func (rectangle *Rectangle) calculatePerimeter() float32 {
	return 2*rectangle.width + 2*rectangle.height
}

// Circle
type Circle struct {
	radius float32
}

func (circle *Circle) calculateArea() float32 {
	return math.Pi * (circle.radius * circle.radius)
}

func (circle *Circle) calculatePerimeter() float32 {
	return 2 * math.Pi * circle.radius
}
