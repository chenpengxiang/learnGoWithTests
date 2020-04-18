package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (rectangle Rectangle) Area() (area float64) {
	return rectangle.Width * rectangle.Height
}


type Circle struct {
	Radius float64
}

func (c Circle) Area() (area float64) {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base float64
	Height float64
}

func (triangle Triangle) Area() (area float64) {
	return (triangle.Base * triangle.Height) * 0.5
}