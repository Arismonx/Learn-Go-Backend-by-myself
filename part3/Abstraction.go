package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}
