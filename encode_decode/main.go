package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// rune: unicode code point
	// start, stop := 'A', 'Z' // rune literals
	var start, stop int
	args := os.Args[1:]
	if len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	if start == 0 || stop == 0 {
		start = 'A'
		stop = 'Z'
		// rune literals are typeless, They can be assigned to a variable
		// with any numeric type
		// each gets converted to int here
	}

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n", "literal", "dec", "hex",
		"encoded", strings.Repeat("-", 45))
	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n))
		// % x with a space prints each byte separated with spaces
		// string(n) encodes the given code point to a utf-8 encoded string automatically
	}

	// var mychar byte = '➾' // doesn't fit into one byte, declare it as rune

	// summary:
	// you can respresent any unicode code point using the rune type
	// because it can store 4 bytes of data

}
