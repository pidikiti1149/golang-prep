package main

import (
	"fmt"
	"math/rand"
)

func main() {

	/*guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 2)
		fmt.Printf("%d ", n)
	}
	fmt.Println()*/

	guess := 10

	for turn := 0; turn < 5; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("You won on first turn")
			return
		} else if n == guess {
			fmt.Println("You won")
			return
		}

	}
	fmt.Println("You Lost")

}
