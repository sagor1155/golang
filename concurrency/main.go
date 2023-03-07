package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Go Concurrency")
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		func() {
			defer wg.Done()
			go work(i)
		}()
	}
	wg.Wait()
	// time.Sleep(time.Second)
	fmt.Println("main is done")
}

func work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task", id, "is done")
}
