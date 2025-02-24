package main

import (
	"fmt"
)

type Drink interface {
	GetName() string
	GetPrice() float32
}

type SelectDrink struct {
	Name  string
	Price float32
}

func (s SelectDrink) GetName() string {
	return s.Name
}

func (s SelectDrink) GetPrice() float32 {
	return s.Price
}

// OCP + LSP
type Coffee struct {
	SelectDrink
}

type Tea struct {
	SelectDrink
}

type Juice struct {
	SelectDrink
}

// DIP
func NewDrink(name string, price float32) Drink {
	return SelectDrink{Name: name, Price: price}

}

func main() {

	coffee := Coffee{SelectDrink{Price: 50, Name: "Coffee"}}
	tea := Coffee{SelectDrink{Price: 50, Name: "Tea"}}
	juice := Coffee{SelectDrink{Price: 50, Name: "Juice"}}

	menu := []Drink{coffee, tea, juice}
	water := NewDrink("Water", 10)
	menu = append(menu, water)

	for _, drink := range menu {
		fmt.Println(drink.GetName(), "Price ", drink.GetPrice(), "Bath")
	}

}
