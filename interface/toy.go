package main

import (
	"fmt"
)

type toy struct {
	title string
	price float64
}

func (t *toy) print() {
	fmt.Printf("%-15s: $%.2f\n", t.title, t.price)
}

func (t *toy) discount(ratio float64) {
	t.price *= float64(1 - ratio)
}
