package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Likes []string
}

func main() {
	fmt.Println("map in go")
	people := [...]*Person{
		{"aa", []string{"apple", "mango"}},
		{"bb", []string{"mango", "banana"}},
	}

	// likemap := make(map[string][]*Person)
	likemap := map[string][]*Person{} // same as above

	fmt.Println(len(people))
	for _, p := range people {
		for _, l := range p.Likes {
			likemap[l] = append(likemap[l], p)
			// fmt.Println(l)
		}
	}

	for _, p := range likemap["mango"] {
		fmt.Println(p.Name, "likes mango.")
	}

	for _, p := range likemap["apple"] {
		fmt.Println(p.Name, "likes apple.")
	}

	for _, p := range likemap["banana"] {
		fmt.Println(p.Name, "likes banana.")
	}

	delete(likemap, "mango")

	for key, value := range likemap {
		for _, p := range value {
			fmt.Println("Key:", key, "Value:", p.Name)
		}
	}
}
