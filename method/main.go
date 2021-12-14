package main

import (
	"fmt"
)

type money float64

func (m money) string() string {
	return fmt.Sprintf("%.2f", m)
}

type game struct {
	title string
	price money
}

func (g game) print() {
	fmt.Printf("%-15s: $%s\n", g.title, g.price.string())
}

type list []*game

func (l list) print() {
	for _, g := range l {
		g.print()
	}
}

func main() {
	var minecraft = game{title: "minecraft", price: 10.0}
	var tetris = game{title: "tetris", price: 20.0}
	var items = []*game{&minecraft, &tetris}
	gameList := list(items)
	gameList.print()
}
