package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const MaxTurn = 5

func main() {
	fmt.Println("Lucky Number")
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Guess number required!!!")
		os.Exit(1)
	}
	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Not a number!!!")
		os.Exit(1)
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		os.Exit(1)
	}

	// dynamically adjust difficulty
	for turn := MaxTurn + guess/4; turn > 0; turn-- {
		n := rand.Intn(guess + 1)
		fmt.Printf("%d ", n)

		if n == guess {
			fmt.Println("\n*** YOU WIN! ***")
			return
		}
	}

	fmt.Println("\n*** YOU LOST! ***")
}
