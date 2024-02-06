// package main

// import "fmt"

// type Shape interface {
//  Area() float64
// }

// type Circle struct {
//  radius float64
// }

// func (c Circle) Area() float64 {
//  return 3.14 * c.radius * c.radius
// }

// type Rectangle struct {
//  Widht  float64
//  Height float64
// }

// func (r Rectangle) Area() float64 {
//  return r.Widht * r.Height
// }

// func pshd(s Shape) {

//  fmt.Printf("Area %f.\n", s.Area())

// }
// func main() {

//  r := Rectangle{10, 12}
//  pshd(r)
//  c := Circle{3}
//  pshd(c)

// }