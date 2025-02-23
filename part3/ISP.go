package main

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type PrinterMachine struct{}

func (p PrinterMachine) Print() {
	fmt.Println("Printing document...")
}
