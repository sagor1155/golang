package main

import "fmt"

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubik", price: 5}
		yoda      = toy{title: "yoda", price: 150}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik, &yoda)
	store.print()
	fmt.Println("--------------------------")

	// interface values are comparable
	// fmt.Println(store[0] == &minecraft)
	// fmt.Println(store[3] == rubik)

	//type assertion: used to extract dynamic value from interface value
	//type assertion: used to check whether an interface value provides the method that you want
	/*
		var p printer
		p = &minecraft // minecraft is dynamic value here, as we can change the value of p in runtime
		p = &tetris    // tetris is dynamic value here, as we can change the value of p in runtime
		p.print()
		tetris.discount(0.5)
		// p.discount(0.5)	// p can't call discount() as it's not declared in printer interface
		p.print()

		p.(*game).discount(0.5) // extracted dynamic value from interface value p
		p.print()
	*/
	// discount all the game in store
	store.discount(0.5)
	store.print()
}
