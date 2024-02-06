package main

import ("fmt"
)

type Shape interface {
	Area() float64
}

type rectangle struct {
	Height float64
	Width  float64
}

func (a rectangle) Area() float64 {
	return a.Height * a.Width
}

type Circle struct {
	Radius float64
}

func (d Circle) Area() float64 {
	return 3.14 * d.Radius * d.Radius
}

func abs(s Shape) {
	fmt.Printf("Area %f.\n", (math.roundFloat(s.Area(), 2)))
}

func main() {
	r := rectangle{5, 3}
	abs(r)
	c := Circle{5}
	abs(c)
}
