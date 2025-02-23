package main

import "fmt"

type Transport interface {
	Drive()
}
type Car struct {
	Name string
}

type ElectricCar struct {
	Car
}

//------------------------------------

type Users struct {
	ID   uint
	Name string
}

func (c Car) Drive() {
	fmt.Println("Name:", c.Name)
}

func main() {
	c := Car{Name: "toyota"}
	e := ElectricCar{Car{Name: "BMW"}}
	c.Drive()
	e.Drive()
}
