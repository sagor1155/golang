package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Testing strings and runes")

	// unicode code points
	// var start, stop int
	// start = 50450
	// stop = 50470
	// fmt.Printf("%-10s %-10s %-10s %-12s\n", "literals", "dec", "hex", "encoded")
	// for n := start; n <= stop; n++ {
	// 	fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n))
	// }

	// var mychar rune = 'ū'
	// fmt.Printf("%-10c % 10x\n", mychar, string(mychar))

	// string to byte slice conversion
	// str := "Yūgen ☯ 💀"
	// fmt.Printf("%s\n", str)
	// fmt.Printf("%d bytes\n", len(str))                    // number of bytes in string
	// fmt.Printf("%d runes\n", utf8.RuneCountInString(str)) // number of runes

	// sbytes := []byte(str) // creates new backing array
	// fmt.Printf("%s\n", sbytes)
	// fmt.Printf("%d bytes\n", len(sbytes))
	// fmt.Printf("%d runes\n", utf8.RuneCount(sbytes))

	// sbytes[0] = 'N'
	// sbytes[1] = 'Y'
	// sbytes[2] = 'C'
	// fmt.Printf("%s\n", sbytes)
	// fmt.Printf("%s\n", str)

	// str = string(sbytes) // creates new backing array
	// fmt.Printf("%s\n", str)

	// indexing and slicing string
	// str := "Yūgen ☯ 💀"
	// for i, r := range str {
	// 	fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	// }
	// fmt.Println()
	// fmt.Printf("1st byte   : %c\n", str[0])
	// fmt.Printf("2nd byte   : %c\n", str[1])   // indexing a string returns a byte value
	// fmt.Printf("2nd rune   : %s\n", str[1:3]) // slicing a string also returns a string value
	// fmt.Printf("last rune  : %s\n", str[len(str)-4:])

	// decode string
	str := "Yūgen ☯ 💀"
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c\n", r)
		i += size
	}

	// modify string
	word := []byte("abcdef")
	fmt.Printf("%s = % [1]x\n", word)
	_, size := utf8.DecodeRune(word)
	copy(word[:size], bytes.ToUpper(word[:size]))
	fmt.Printf("%s = % [1]x\n", word)

	// string header
	// why strings are immutable?
}
