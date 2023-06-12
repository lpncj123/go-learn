package main

import "fmt"

type Shape interface {
	area() float64
}
type Rectangle struct {
	height float64
	width  float64
}
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * 3.14
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func main() {
	var s Shape
	//b:=new(Rectangle)
	s = Rectangle{1, 2}
	fmt.Printf("%f", s.area())
	s = Circle{2}
	fmt.Printf("%f", s.area())
}
