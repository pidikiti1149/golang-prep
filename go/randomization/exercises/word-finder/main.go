package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := strings.Fields("Hello aditya aditya ela unav nuvu")

	query := os.Args[1:]

	for _, q := range query {

		for i, w := range words {

			if strings.ToLower(q) == strings.ToLower(w) {

				fmt.Printf("#%-2d: %q\n", i+1, w)
				break
			}
		}
	}

}
