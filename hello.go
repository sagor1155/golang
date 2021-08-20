package main

import "fmt"

/*
GO:
1. array
2. Dynamic memory - allocation/deallocation
3. OOP (class, object, encapsulation polymorphism, inheritence, abstraction, interface)
4. static
5. Pointer
6. slice
7. maps
8. exception handling
9. File I/O
10. NULL pointer
11. struct
12. typedef
13. garbage collector
14. Design pattern
*/

func calc(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}

func main() {
	fmt.Println("Hello! World")

	// variables
	var x int8 = 123
	var y = 456
	z := 67
	z = 90
	var m float32 = 234.55
	const n = 10
	fmt.Println("x: ", x, ", y: ", y, " z: ", z, " m: ", m, ", n:", n)

	// for loop
	var i int
	for i = 0; i < 10; i++ { // var declaration not allowed inside for loop
		if i >= 5 && i <= 9 {
			fmt.Print(i, " ")
		} else { // else must be on this line
			fmt.Print(i+5, " ")
		}
	}
	fmt.Println("")

	// switch-case
	a, b := 2, 1
	switch (a / 5) + (b * 2) {
	case 1:
		fmt.Println("Sum is 1")
	case 2:
		fmt.Println("Sum is 2")
	case 3:
		fmt.Println("Sum is 3")
	default:
		fmt.Println("Printing Default")
	}

	// array
	// arr := [3]string{"one", "two", "three"}
	var arr [3]string
	arr[0] = "one"
	arr[1] = "two"
	arr[2] = "three"
	fmt.Println(arr[0], arr[1], arr[2])

	directions := [...]int{1, 2, 3, 4, 5}
	for i = 0; i < len(directions); i++ {
		fmt.Println(directions[i])
	}

	// slice
	aa := [5]string{"1", "2", "3", "4", "5"}
	slice_a := aa[1:4]
	bb := [5]string{"one", "two", "three", "four", "five"}
	slice_b := bb[1:4]

	fmt.Println("Slice_a:", slice_a)
	fmt.Println("Slice_b:", slice_b)
	fmt.Println("Length of slice_a:", len(slice_a))
	fmt.Println("Length of slice_b:", len(slice_b))

	slice_a = append(slice_a, slice_b...) // appending slice
	fmt.Println("New Slice_a after appending slice_b :", slice_a)

	slice_a = append(slice_a, "text1") // appending value
	fmt.Println("New Slice_a after appending text1 :", slice_a)
	slice_a[0] = "werw"
	fmt.Println(aa)
	fmt.Println(slice_a)

	// function
	fmt.Println(calc(5, 8))

	xx, yy := 15, 10
	sum, diff := calc(xx, yy)
	fmt.Println("Sum", sum)
	fmt.Println("Diff", diff)
}
