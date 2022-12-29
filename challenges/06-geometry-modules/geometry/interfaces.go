package geometry

import (
	"fmt"
)

/*
Interface
- name in lowercase to do it private -
*/
type geometry interface {
	calculateArea() float32
	calculatePerimeter() float32
}

func PrintGeometricMeasures(g geometry) {
	fmt.Println("Area: ", g.calculateArea())
	fmt.Println("Perimeter: ", g.calculatePerimeter())
}
