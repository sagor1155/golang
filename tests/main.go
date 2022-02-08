package main

import (
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
	str := "Yūgen ☯ 💀"
	fmt.Printf("%s\n", str)
	fmt.Printf("%d bytes\n", len(str))                    // number of bytes in string
	fmt.Printf("%d runes\n", utf8.RuneCountInString(str)) // number of runes

	sbytes := []byte(str) // creates new backing array
	fmt.Printf("%s\n", sbytes)
	fmt.Printf("%d bytes\n", len(sbytes))
	fmt.Printf("%d runes\n", utf8.RuneCount(sbytes))

	sbytes[0] = 'N'
	sbytes[1] = 'Y'
	sbytes[2] = 'C'
	fmt.Printf("%s\n", sbytes)
	fmt.Printf("%s\n", str)

	str = string(sbytes) // creates new backing array
	fmt.Printf("%s\n", str)

}
