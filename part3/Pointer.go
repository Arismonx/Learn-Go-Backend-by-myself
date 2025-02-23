package main

import "fmt"

func Pointer(num *int) {
	var G *int
	G = num
	*G = *num + 100
	fmt.Println("G: ", *G)
}
