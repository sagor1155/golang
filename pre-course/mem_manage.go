package main

import (
	"fmt"
)

// we can return address of local variable in GO, but not in C/C++
// as GO keeps these variable in heap when local variable address is returned from function
// 'x' escapes to the heap (at compile time not in runtime)
// sharing down typically stays on the stack
// sharing up typically stays on the heap
func answere() *int {
	x := 42
	return &x // as we are returning address of variable x, compiler knows it's
	// not safe to keep variable x in stack, so it keeps x in heap
}

func main() {
	fmt.Println("Memory management in GO")
	n := answere()
	fmt.Println(*n / 2)
}
