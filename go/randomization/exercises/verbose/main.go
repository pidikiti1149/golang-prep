package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	/*guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 2)
		fmt.Printf("%d ", n)
	}
	fmt.Println()*/

	args := os.Args

	if len(args) != 3 {
		fmt.Println("Give me two number")
		return
	}

	ver := args[1]

	var guess int

	guess2, err := strconv.Atoi(args[2])

	if guess2 <= 10 {
		guess = 10
	} else {
		guess = guess2
	}

	if err != nil {
		fmt.Println("Give two integer")
		return
	}

	var m int

	switch {
	case guess <= 10:
		m = 5
	case guess <= 25:
		m = 10
	case guess <= 50:
		m = 15
	case guess <= 75:
		m = 20
	case guess <= 100:
		m = 25
	case guess > 100:
		m = 30
	}
	for turn := 0; turn < m; turn++ {

		n := rand.Intn(guess + 1)

		if ver == "-v" {
			fmt.Printf(" %d", n)
		}

		if n == guess2 {
			fmt.Println(" You won")
			return
		}
	}
	fmt.Println(" You Lost")

}
