package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "is making a sound")
}

type Dog struct {
	Animal // Embedded struct (เหมือนสืบทอด)
}
