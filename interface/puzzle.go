package main

import (
	"fmt"
)

type puzzle struct {
	title string
	price float64
}

func (p puzzle) print() {
	fmt.Printf("%-15s: $%.2f\n", p.title, p.price)
}
