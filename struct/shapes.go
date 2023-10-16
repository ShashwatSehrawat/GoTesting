package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
