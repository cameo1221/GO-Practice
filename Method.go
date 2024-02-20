package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	r := Rectangle{Width: 4.5, Height: 50.5}
	c := Circle{Radius: 39.7}
	fmt.Printf("The area of the rectangle is : %.2f\n ", r.Area())
	fmt.Printf("The area of the circle is %.2f\n", c.Area())
}
