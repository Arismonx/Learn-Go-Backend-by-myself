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

	fmt.Println("--------------------")
	//----------  Pointer ---------------
	// var Tuschy *int
	// Lotus := 500
	// var Nami *int

	// Tuschy = &Lotus

	// fmt.Println("Tushy: ", *Tuschy)
	// fmt.Println("Lotus: ", Lotus)

	// Pointer(Tuschy)

	// fmt.Println("Tushy: ", *Tuschy)
	// fmt.Println("Lotus: ", Lotus)

	// Nami = &Lotus
	// fmt.Println("Nami: ", *Nami)

	// *Nami = *Nami - 400
	// fmt.Println("Tushy: ", *Tuschy)
	// fmt.Println("Lotus: ", Lotus)
	// fmt.Println("Nami: ", *Nami)

	// Pointer(Nami)
	// fmt.Println("Tushy: ", *Tuschy)
	// fmt.Println("Lotus: ", Lotus)
	// fmt.Println("Nami: ", *Nami)

	//----------S - Single Responsibility Principle (SRP)---------------

	r := Report{Title: "Hello Everyone!"}
	SaveReport(r)

	//----------O - Open/Closed Principle (OCP)---------------

	rec := Rectangle_1{width: 10, height: 2}
	circle_1 := Circle_1{radius: 2}

	PrintArea_1("Rectangle", rec)
	PrintArea_1("Circle", circle_1)

	//----------L - Liskov Substitution Principle (LSP)---------------
	spa := Sparrow{}

	spa.Fly()
	makeBirdFly(spa)

	//---------- I - Interface Segregation Principle (ISP)---------------

	p := PrinterMachine{}
	p.Print()

	//---------- D - Dependency Inversion Principle (DIP) ---------------

	email := EmailNotifier{}
	service := NotificationService{notifier: email}
	service.NotifyUser("Hello!")
	email.SendNotification("Tuschy")
}
