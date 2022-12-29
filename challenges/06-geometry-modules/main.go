package main

import (
	"fmt"

	"github.com/AntonioMartinezFernandez/udemy-golang-profesional/challenges/06-geometry-modules/geometry"
)

func main() {
	rectangle := geometry.Rectangle{
		Width:  5.2,
		Height: 10.6,
	}

	circle := geometry.Circle{
		Radius: 8.1,
	}

	fmt.Println("RECTANGLE", rectangle)
	geometry.PrintGeometricMeasures(&rectangle)

	fmt.Println("CIRCLE", circle)
	geometry.PrintGeometricMeasures(&circle)
}
