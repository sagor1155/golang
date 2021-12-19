package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	// var names []byte

	// optimize for huge number of files in a directory
	// pre allocate names slice
	// total := len(files) * 256 // number of files * average max file name size = 255 + 1 for '\n'
	// names := make([]byte, 0, total)
	// fmt.Printf("Total allocated space: %d bytes.\n", total)

	// get the exact size, more optimized in terms of memory
	var total_exact int
	for _, file := range files {
		if file.Size() == 0 {
			total_exact += len(file.Name()) + 1
		}
	}
	fmt.Printf("Total required space: %d bytes.\n", total_exact)
	names := make([]byte, 0, total_exact)

	// empty file finder
	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			// fmt.Println(name, file.Size())
			names = append(names, name...)
			names = append(names, '\n')
			// string value is actually a read-only byte slice. The ellipsis passes the
			// byte elements of the name string as individual elements to the append function
		}
	}
	err = ioutil.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", names)
}
