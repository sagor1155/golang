package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	s "github.com/inancgumus/prettyslice"
	"github.com/inancgumus/screen"
)

func slicePracticeBasics() {
	fmt.Println("Example for Slices: Dynamic array in GO")
	/*
		array is static (can't grow and shrink)
		slices are dynamic (can grow and shrink)
	*/

	trees := []string{} //empty, len=0
	fmt.Printf("%#v\n", trees)
	trees = append(trees, "pine", "coconut") //append returns new slice
	fmt.Printf("%#v\n", trees)
	fmt.Println(len(trees))
	fmt.Println(trees[0])

	flowers := []string{"marrygold", "rose"}
	trees = append(trees, flowers...)
	fmt.Printf("%#v\n", trees)
	fmt.Printf("%T\n", trees)

	var books []string //nil, len=0
	fmt.Printf("%#v\n", books)
	fmt.Println(len(books))
	// fmt.Println(books[0]) // gives error

	//sliceable: array, slice, string (as it's a slice of bytes)
	//sliceable[START:STOP]
	//START - index (starts from 0) - slice it starting from this
	//STOP - index+1 (starts from 1) - slice it upto this
	msg := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf("%#v\n", msg)
	newMsg := msg[2:]
	fmt.Printf("%#v\n", newMsg)

	mymsg := append(msg[:4], '!') //actual input slice 'msg' will also be updated??? note that
	fmt.Printf("%#v\n", msg)
	fmt.Printf("%#v\n", mymsg)

	//pagination example
	numbers := []int{1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13}

	const pageSize = 4
	l := len(numbers)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}
		currentPage := numbers[from:to]
		head := fmt.Sprintf("page #%d", (from / pageSize))
		fmt.Printf("%s  %v\n", head, currentPage)
	}

	agesArray := [3]int{1, 2, 3} // backing array for slice ages
	ages := agesArray[0:3]       // shares the same memory as agesArray
	ages1 := ages[0:1]
	ages2 := ages[1:3]

	fmt.Printf("agesArray %v\n", agesArray)
	fmt.Printf("ages %v\n", ages)
	fmt.Printf("ages1 %v\n", ages1)
	fmt.Printf("ages2 %v\n", ages2)

	ages[0] = 33
	ages[2] = 77

	fmt.Println("---------------------------")
	fmt.Printf("agesArray %v\n", agesArray)
	fmt.Printf("ages %v\n", ages)
	fmt.Printf("ages1 %v\n", ages1)
	fmt.Printf("ages2 %v\n", ages2)

	fmt.Printf("addr agesArray %d\n", &agesArray[0])
	fmt.Printf("addr ages %d\n", &ages[0])
}

func testAppendBehaviour() {
	s.PrintBacking = true
	s.MaxPerLine = 30
	s.Width = 150

	var nums []int

	screen.Clear()
	for cap(nums) <= 128 {
		screen.MoveTopLeft()
		s.Show("nums", nums)
		nums = append(nums, rand.Intn(9)+1)
		time.Sleep(time.Second / 4)
	}
}

func testAppendGrowthRate() {
	ages, oldCap := []int{1}, 1.0

	for len(ages) < 5e5 {
		ages = append(ages, 1)
		c := float64(cap(ages))
		if c != oldCap {
			fmt.Printf("len:%-10d cap:%-10g growth:%.2f\n", len(ages), c, c/oldCap)
		}
		oldCap = c
	}
}

func testComparable() {
	// arrays
	books := [3]string{"dracula", "1984", "island"}
	newBooks := [3]string{"ulysses", "fire"}
	books = newBooks // can't be assigned to each other if array type (element type + length) is not same
	fmt.Println(books)
}

func stringReplaceExample() {
	// string replace
	str := "all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay"
	lyric := strings.Split(str, " ")
	fmt.Printf("%#v\n", lyric)
	lyric = append([]string{"yesterday"}, lyric...)
	fmt.Println(lyric)
	const N, M = 8, 13
	lyric = append(lyric, lyric[N:M]...)
	fmt.Println(lyric)
	lyric = append(lyric[:N], lyric[M:]...)
	fmt.Println(lyric)
}

func testBackingArray() {
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

func testSliceHeader() {
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

func testMultiDimensionalSlice() {
	// multi-dimensional slice
	// daily spendings
	// 1st day: $200, $100
	// 2nd day: $500
	// 3rd day: $50, $25, $75

	// spendings := [][]int{
	// 	{200, 100},
	// 	{500},
	// 	{50, 25, 75},
	// }

	// spendings := make([][]int, 0, 5)
	// spendings = append(spendings, []int{200, 100})
	// spendings = append(spendings, []int{25, 10, 45, 60})
	// spendings = append(spendings, []int{5, 15, 35})
	// spendings = append(spendings, []int{95, 10}, []int{50, 25})

	spendings := fetch()
	for i, daily := range spendings {
		// fmt.Println(i+1, daily)
		var total int
		for _, spending := range daily {
			total += spending
		}
		fmt.Printf("Day %d: $%d\n", i+1, total)
	}
}

func fetch() [][]int {
	content := `200 100
25 10 45 60
5 15 35
95 10
50 25`
	lines := strings.Split(content, "\n")
	spendings := make([][]int, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		spendings[i] = make([]int, len(fields))
		for j, field := range fields {
			spending, _ := strconv.Atoi(field)
			spendings[i][j] = spending
		}
	}
	return spendings
}

func main() {
	// testAppendBehaviour()
	// testAppendGrowthRate()
	// testComparable()
	// stringReplaceExample()
	// testBackingArray()
	// testSliceHeader()
	testMultiDimensionalSlice()
}
