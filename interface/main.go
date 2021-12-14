package main

import (
	"fmt"
	"strconv"
)

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

	// empty interface
	// means nothing, any type of value can be stored in it
	// but can't directly use the dynamic value of an empty interface value in operations
	var any interface{}
	any = []int{1, 2, 3}
	any = map[int]bool{1: true, 2: false}
	any = 3 // the dynamic type is int here
	// any = any * 2 // invalid operation because any is an empty interface
	// to do this, need to extract dymanic value using type assertion
	any = any.(int) * 2
	fmt.Println(any)

	nums := []int{1, 2, 3}
	any = nums // dynamic type of any is []int and value is []int{1, 2, 3}
	// we can't use 'any' as an int slice. first we need to extract it
	// _ = len(any) // invalid operation
	fmt.Println(len(any.([]int))) // extract dynamic value using type assertion
	// program will crash here if the dynamic type is not an int slice

	any = append(any.([]int), 2)
	fmt.Println(any)
	fmt.Println(nums)

	// slice of empty interface value
	var many []interface{}
	// many = nums // error: many is a slice of empty interface it's not just interface, so can't put nums here
	// rather we can append some empty interfaces or, other types
	for _, n := range nums {
		many = append(many, n)
	}
	fmt.Println(many)

	fmt.Println(format(345))
	fmt.Println(format("776"))
}

func format(v interface{}) int {
	var t int

	if v, ok := v.(int); ok {
		t = v
	}

	if v, ok := v.(string); ok {
		t, _ = strconv.Atoi(v)
	}

	return t
}
