package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var i = 10

func short_declaration() {
	// short declaration
	width, height := 100, 50
	width, color := 50, "red"
	fmt.Println(width, height, color)
}

func shadowing_variable() {
	// shadowing variable
	fmt.Println("package's i:", i)

	// package's i is being shadowed:
	i := 5
	// i above is a new variable.

	fmt.Println("main block:", i)

	// a new scope begins
	{
		// i is a new variable
		i := 6
		fmt.Println("inner block:", i)
	}
	// the scope ends

	// main's i
	fmt.Println("main's i:", i)
}

func type_conversion() {
	speed := 100
	force := 2.5

	speed = int(float64(speed) * force)
	fmt.Println(speed)
}

func commandline_arguments() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println(len(os.Args))
	for _, val := range os.Args {
		fmt.Println(val)
	}
}

func string_length() {
	name := "inglés"
	// length of string
	fmt.Println("string len:", len(name))
	// number of character in string
	fmt.Println("number of character:", utf8.RuneCountInString(name))

	name = " english "
	s := "    " + strings.Repeat("!", len(name)) + strings.ToUpper(name) + strings.Repeat("!", len(name)) + "    "
	fmt.Println(s)
	fmt.Println(strings.TrimSpace(s))
}

func constant_and_iota() {
	const (
		EST = -(5 + iota) // iota is a constant number generator used within const, increases by one, same expression for all the items below
		_                 // blank identifier to skip
		MST
		PST
	)
	fmt.Println(EST, MST, PST)
}

func main() {
	// short_declaration()
	// shadowing_variable()
	// type_conversion()
	// commandline_arguments()
	// string_length()
	constant_and_iota()

}
