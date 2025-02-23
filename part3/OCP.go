package main

import "fmt"

type Shape_1 interface {
	Area_1() float64
}

type Rectangle_1 struct {
	width, height float64
}

func (r Rectangle_1) Area_1() float64 {
	return r.width * r.height
}

type Circle_1 struct {
	radius float64
}

func (c Circle_1) Area_1() float64 {
	return 3.14 * c.radius * c.radius
}

func PrintArea_1(str string, s Shape_1) {
	fmt.Printf("%s of Area: %v \n", str, s.Area_1())
}
