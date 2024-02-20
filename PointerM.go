package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func Area(r rectangle) {
	fmt.Printf("Area of length %d and width %d is: %d\n", (r.length), (r.width), (r.length * r.width))
}

func (r rectangle) Area() {
	fmt.Printf("Area of length %d and width %d is: %d\n", (r.length), (r.width), (r.length * r.width))
}

func main() {
	r := rectangle{
		length: 50,
		width:  12,
	}
	Area(r)
	r.Area()
	p := &r
	p.Area()
}
