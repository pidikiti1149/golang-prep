package main

import (
	"fmt"
)

var max = 5

var i int

func main() {

	fmt.Printf("%5s", "X")

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {

			fmt.Printf("%5d", i*j)

		}

		fmt.Println()
	}
}
