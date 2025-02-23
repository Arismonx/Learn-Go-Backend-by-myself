package main

import (
	"fmt"
)

func main() {
	//----------Encapsulation---------------
	user := User{}
	user.SetName("Tuschy")
	user.Age = 22
	user.name = "New Tuschy" // Same Package can call user.name
	fmt.Println(user.GetName(), user.Age)

	//----------Abstraction---------------
	circle := Circle{radius: 5}
	printArea(circle)

	circle.radius = 4
	circle.Area()
	printArea(circle)

	//----------Inheritance ---------------

	d := Dog{Animal{Name: "Buddy"}}
	d.Speak() // Buddy is making a sound

	//---------- Polymorphism  ---------------
	obj_dog := Dog2{}
	obj_cat := Cat{}

	makeSound(obj_dog)
	makeSound(obj_cat)

	//----------  Pointer ---------------
	var Tuschy *int
	Lotus := 500
	var Nami *int

	Tuschy = &Lotus

	fmt.Println("Tushy: ", *Tuschy)
	fmt.Println("Lotus: ", Lotus)

	Pointer(Tuschy)

	fmt.Println("Tushy: ", *Tuschy)
	fmt.Println("Lotus: ", Lotus)

	Nami = &Lotus
	fmt.Println("Nami: ", *Nami)

	*Nami = *Nami - 400
	fmt.Println("Tushy: ", *Tuschy)
	fmt.Println("Lotus: ", Lotus)
	fmt.Println("Nami: ", *Nami)

	Pointer(Nami)
	fmt.Println("Tushy: ", *Tuschy)
	fmt.Println("Lotus: ", Lotus)
	fmt.Println("Nami: ", *Nami)
}
