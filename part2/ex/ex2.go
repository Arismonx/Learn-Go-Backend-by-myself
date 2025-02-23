package main

import (
	"fmt"
	"time"
)

type Empoyee struct {
	Name     string
	Position string
	Salary   float32
}

func greet(name string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Name :", name)
}

func sendOk(message chan<- string) {
	time.Sleep(2 * time.Second)
	message <- "Send Data OK!"
}

func main() {
	start := time.Now()
	ch := make(chan string)
	var emp Empoyee
	var emp1 Empoyee
	var emp2 Empoyee

	emp.Name = "Tuschy"
	emp1.Name = "Nrinee"
	emp2.Name = "Mina"
	// emp.Position = "Deverloper"
	// emp.Salary = 50000.20

	go greet(emp.Name)
	go greet(emp2.Name)
	go sendOk(ch)
	greet(emp1.Name)

	if GetData := <-ch; GetData == "Send Data OK!" {
		fmt.Println("GET IT COMPLETE!")
		fmt.Println("Time :", time.Since(start), "Sce.")
	}
}
