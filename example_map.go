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
	people := [...]*Person{{"aa", []string{"apple", "mango"}}, {"bb", []string{"mango", "banana"}}}

	// likes := make(map[string][]*Person)
	likes := map[string][]*Person{} // same as above

	fmt.Println(len(people))
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
			// fmt.Println(l)
		}
	}

	for _, p := range likes["mango"] {
		fmt.Println(p.Name, "likes mango.")
	}

	for _, p := range likes["apple"] {
		fmt.Println(p.Name, "likes apple.")
	}

	for _, p := range likes["banana"] {
		fmt.Println(p.Name, "likes banana.")
	}

	delete(likes, "mango")

	for key, value := range likes {
		for _, p := range value {
			fmt.Println("Key:", key, "Value:", p.Name)
		}
	}
}
