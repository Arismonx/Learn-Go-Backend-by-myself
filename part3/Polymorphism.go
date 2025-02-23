package main

import "fmt"

type Speaker interface {
	Speak()
}

type Cat struct{}
type Dog2 struct{}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

func (d Dog2) Speak() {
	fmt.Println("Woof")
}

func makeSound(s Speaker) {
	s.Speak()
}
