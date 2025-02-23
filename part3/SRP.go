package main

import "fmt"

type Report struct {
	Title string
}

func SaveReport(r Report) {
	fmt.Println(r.Title)
}
