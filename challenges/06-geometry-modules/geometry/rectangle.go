package geometry

// Rectangle Class
type Rectangle struct {
	Width  float32
	Height float32
}

func (rectangle *Rectangle) calculateArea() float32 {
	return rectangle.Width * rectangle.Height
}

func (rectangle *Rectangle) calculatePerimeter() float32 {
	return 2*rectangle.Width + 2*rectangle.Height
}
