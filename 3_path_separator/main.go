package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	dir, file = path.Split("/home/sagor/work/main.css")
	fmt.Println("Directory:", dir)
	fmt.Println("file:", file)
}
