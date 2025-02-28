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

	var guess int

	if len(args) != 3 {
		fmt.Println("Give me two number")
		return
	}

	guess1, err := strconv.Atoi(args[1])

	guess2, err1 := strconv.Atoi(args[2])

	if guess1 > guess2 {
		guess = guess1
	} else {
		guess = guess2
	}

	if err != nil {
		fmt.Println("Give two integer")
		return
	} else if err1 != nil {
		fmt.Println("Give two integer")
		return
	}

	for turn := 0; turn < 5; turn++ {

		n := rand.Intn(guess + 1)

		if n == guess {

			switch rand.Intn(3) {

			case 0:
				fmt.Println("You won")
				return
			case 1:
				fmt.Println("You're  Awesome")
				return
			case 2:
				fmt.Println("You're  Perfect")
				return
			}
		}

	}

	switch rand.Intn(2) {

	case 0:
		fmt.Println("You Lost")
		return
	case 1:
		fmt.Println("You're a Looser")
		return
	}

}
