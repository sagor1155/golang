package main

import (
	"fmt"
)

func main() {
	fmt.Println("Example for Slices: Dynamic array in GO")
	/*
		array is static (can't grow and shrink)
		slices are dynamic (can grow and shrink)
	*/

	trees := []string{} //empty, len=0
	fmt.Printf("%#v\n", trees)
	trees = append(trees, "pine", "coconut") //append returns new slice
	fmt.Printf("%#v\n", trees)
	fmt.Println(len(trees))
	fmt.Println(trees[0])

	flowers := []string{"marrygold", "rose"}
	trees = append(trees, flowers...)
	fmt.Printf("%#v\n", trees)
	fmt.Printf("%T\n", trees)

	var books []string //nil, len=0
	fmt.Printf("%#v\n", books)
	fmt.Println(len(books))
	// fmt.Println(books[0]) // gives error

	//sliceable: array, slice, string (as it's a slice of bytes)
	//sliceable[START:STOP]
	//START - index (starts from 0) - slice it starting from this
	//STOP - index+1 (starts from 1) - slice it upto this
	msg := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf("%#v\n", msg)
	newMsg := msg[2:]
	fmt.Printf("%#v\n", newMsg)

	mymsg := append(msg[:4], '!') //actual input slcie 'msg' will also be updated??? note that
	fmt.Printf("%#v\n", msg)
	fmt.Printf("%#v\n", mymsg)

	//pagination example
	numbers := []int{1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13}

	const pageSize = 4
	l := len(numbers)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}
		currentPage := numbers[from:to]
		head := fmt.Sprintf("page #%d", (from / pageSize))
		fmt.Printf("%s  %v\n", head, currentPage)
	}
}
