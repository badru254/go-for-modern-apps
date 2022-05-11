package main

import (
	"fmt"
	"math"
)

func main() {

	s := Square{20}
	fmt.Printf("Square of side %v has area of %v\n", s.Length, s.Area())

	c := Circle{20}
	fmt.Printf("Circle of radius %v has area of %.2f\n", c.Radius, c.Area())

	shapes := []Shape{s, c}

	sa := sumAreas(shapes)

	fmt.Printf("Sum of Areas = %.2f\n", sa)

}

//Square is a struct
type Square struct {
	Length float64
}

//Circle is a struct
type Circle struct {
	Radius float64
}

//Area returns the area of the square
func (s Square) Area() float64 {
	return s.Length * s.Length
}

//Area returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// sumAreas returns the sum of all areas in a slice
func sumAreas(shapes []Shape) float64 {

	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

//Shape is a interface
type Shape interface {
	Area() float64
}
