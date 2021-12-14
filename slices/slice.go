package main

import (
	"fmt"
	"sort"

	s "github.com/inancgumus/prettyslice"
)

func backingArray() {
	s.PrintBacking = true
	s.MaxPerLine = 7

	grades := []float64{40, 10, 20, 50, 60, 70}
	s.Show("grades before sort", grades[:])

	front := grades[0:3] // front and grades has same backing array
	sort.Float64s(front)

	s.Show("front", front)
	s.Show("grades after sort", grades[:])

	// to have different backing array
	// var newGrades []float64
	newGrades := append([]float64(nil), grades...)
	s.Show("newGrades", newGrades)
	newGrades[0] = 30

	front2 := newGrades[:3]
	front3 := front2

	sort.Float64s(front2)

	s.Show("front2", front2)
	s.Show("front3", front3)
	s.Show("newGrades", newGrades)
}

type collection []string

func sliceHeader() {
	s.PrintElementAddr = true
	// header: ptr, len, cap
	data := collection{"learn", "slices", "in", "go"}
	change(data) //copies the data slice header to data, but doesn't make copy of backing array
	s.Show("data in calling", data)
	fmt.Printf("addr of data slice in calling: %p\n", &data)
}

func change(data collection) {
	data[1] = "dynamic array"
	s.Show("data in called", data)
	fmt.Printf("addr of data slice in called: %p\n", &data)
}

func main() {
	// backingArray()
	sliceHeader()

}
