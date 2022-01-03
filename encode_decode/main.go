package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func testRuneBasics() {
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
		// with any numeric type, each gets converted to int here
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

func main() {
	// testRuneBasics()
	// let's convert, index and slice bytes, runes and strings
	str := "Yūgen ☯ 💀"
	// str[0] = 'N'
	// str[1] = 'Y'
	// strings are immutable

	sbytes := []byte(str)
	// bytes[0] = 'N'
	// bytes[1] = 'Y'
	//// bytes[2] = 0

	// // convert byte slice to string
	// str = string(bytes)
	// // still they don't share same backing array, because
	// // it creates a string and copies the bytes of the byte slice to the new string's backing array
	// fmt.Printf("%s\n", str) // NY�gen ☯ 💀

	fmt.Printf("%s\n", str)                                 // Yūgen ☯ 💀
	fmt.Printf("\t%d bytes\n", len(str))                    // 15 bytes, len() counts the bytes not the rune
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str)) // counts rune in string, 9 runes
	// 9 runes but 15 bytes because some runes has multiple bytes

	// creates a []byte and copies the bytes of a string to the new slice's backing array
	// changing the byte slice doesn't really change the original str as they don't share the same backing array
	fmt.Printf("% x\n", sbytes)                        // 4e 59 ab 67 65 6e 20 e2 98 af 20 f0 9f 92 80
	fmt.Printf("\t%d bytes\n", len(sbytes))            // 15 bytes, len() counts the bytes not the rune
	fmt.Printf("\t%d runes\n", utf8.RuneCount(sbytes)) // counts rune in byte slice, 9 runes
	// 9 runes but 15 bytes because some runes has multiple bytes

	fmt.Println()
	for i, r := range str {
		// fmt.Printf("str[%2d] = %-2x\n", i, str[i])
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}
	/*
		str[ 0] = 59
		str[ 1] = c5
		str[ 3] = 67
		str[ 4] = 65
		str[ 5] = 6e
		str[ 6] = 20
		str[ 7] = e2
		str[10] = 20
		str[11] = f0

		The indexes don't increase one by one.
		Now after printing the utf encoded version of the rune -

		str[ 0] = 59           = 'Y'
		str[ 1] = c5 ab        = 'ū'
		str[ 3] = 67           = 'g'
		str[ 4] = 65           = 'e'
		str[ 5] = 6e           = 'n'
		str[ 6] = 20           = ' '
		str[ 7] = e2 98 af     = '☯'
		str[10] = 20           = ' '
		str[11] = f0 9f 92 80  = '💀'

		for-range loop jumps overs the runes instead of the bytes of a string.
		Each index returns the starting index of the next runes.
		utf-8 is a variable length encoding, so each rune may start at different index and may have different number of bytes
	*/

	// indexing and slicing a string
	fmt.Println()
	fmt.Printf("1st byte   : %c\n", str[0])
	fmt.Printf("2nd byte   : %c\n", str[1])   // indexing a string returns a byte value
	fmt.Printf("2nd rune   : %s\n", str[1:3]) // slicing a string also returns a string value
	fmt.Printf("last rune  : %s\n", str[len(str)-4:])

	/*
		Runes in a utf-8 encoded string can have a different number of bytes
		because utf-8 is a variable byte-length encoding.

		Especially in scripting language you can index utf-8 string easily.
		However, Go doesn't allow you to do so by deafult beacause of efficiency reasons.
	*/

	// there is a way to index utf-8 string easily
	// convert string to rune slice
	runes := []rune(str) // []rune(string) creates a new slice, and copies each rune to the new slice's backing array
	fmt.Println()
	fmt.Printf("%s\n", string(runes))
	fmt.Printf("\t%d bytes\n", int(unsafe.Sizeof(runes[0]))*len(runes)) // 36 bytes
	fmt.Printf("\t%d runes\n", len(runes))
	fmt.Printf("1st rune   : %c\n", runes[0])
	fmt.Printf("2nd rune   : %c\n", runes[1])  // indexing
	fmt.Printf("first five : %c\n", runes[:5]) // slicing
	// why don't we use rune all the time if indexing is easier?
	// go forces to use string or, byte slice
	// a string value usually use utf-8 so it can be efficient because each rune uses 1-4 bytes (variable byte length)
	// each rune in []rune has the same length: 4 bytes (inefficient), because rune is an alias to int32 (4 byte)

	// string decoding: get all individual runes using index
	// How to get runes from an utf-8 encoded string without using 'for range'
	// extract rune start-end position (index) in a utf-8 string
	const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%c", r)
		i += size
	}
	fmt.Println()
	fmt.Println()

	// get all runes from string using for-range loop
	// for _, r := range text {
	// 	fmt.Printf("%c", r)
	// }
	// fmt.Println()

	// extract and modify rune
	word := []byte("öykü")
	fmt.Printf("%s = % [1]x\n", word)

	var size int
	// for i := range string(word) { // []byte to string conversion is inefficient
	// 	if i > 0 {
	// 		size = i
	// 		break
	// 	}
	// }

	// instead of using for loopp, use the following function, its efficient
	_, size = utf8.DecodeRune(word)

	// modify first rune
	copy(word[:size], bytes.ToUpper(word[:size]))
	fmt.Printf("%s = % [1]x\n", word)
}
