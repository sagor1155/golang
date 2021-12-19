package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Map basics")
	/* map is collection type data
	stores data in key-value paires (Uses Hash table algorithm)
	allows fast look up for the values
	time complexity of fast-lookup for map key is O(1) [constant time]
	Key should be comparable. Slices, maps, and function values are not comparable. So, they cannot be map keys.
	a map value in go is a pointer to a map header
	*/

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Argument missing!")
		os.Exit(1)
	}

	// declaring a map
	// var english map[string]string // nil, len=0
	// english := map[string]string{} // empty, len=0
	// english := make(map[string]string, len) // allocates memory, len is hust a hint, won't actually allocates that size unlike slice

	english := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mukemmel",
	}

	english["up"] = "yukari"

	// map length
	fmt.Println("map length:", len(english))

	// range over a map
	for key, value := range english {
		fmt.Printf("%s: %s\n", key, value)
	}

	// check if key exists
	// value, ok := mapValue[key] // ok is true if key exists, false otherwise
	if _, ok := english[args[0]]; ok {
		fmt.Println("search key exists in english dictionary")
		fmt.Println(args[0], "in turkish is", english[args[0]])
	}

	/* map header:
	map value is only a pointer (contains address) to a map header
	map header contains another pointer (Buckets) to the real data
	passing map to a function or, assigning to a varable is very cheap as we are just passing the pointer
	*/
	dict := english // they are pointing to same map header, changing one map will change the other
	dict["good"] = "guzel"
	english["great"] = "kusursuz"
	fmt.Printf("english: %q\ndict: %q\n", english, dict)

	// make the map
	turkish := make(map[string]string)
	for k, v := range english {
		turkish[v] = k
	}
	fmt.Printf("english: %q\n", english)
	fmt.Printf("turkish: %q\n", turkish)

	// delete key
	delete(english, "up")
	fmt.Printf("english: %q\n", english)

	// delete whole map
	// this is just a single operation behind the scence. go converts the whole loop to a single operation - mapclear()
	for k := range english {
		delete(english, k)
	}

}

/*
type hmap struct {
	count      int
	flags      uint8
	B          uint8
	noverflow  uint16
	hash0      uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr
	extra      *mapextra
}
*/
