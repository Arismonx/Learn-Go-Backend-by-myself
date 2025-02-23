package main

import "fmt"

type Bird interface {
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow flying...")
}

func makeBirdFly(b Bird) {
	b.Fly()
}
