package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	words := filepath.SplitList(os.Getenv("PATH"))

	query := os.Args[1:]
	for _, q := range query {
		for i, w := range words {

			if strings.ContainsAny(w, q) {

				fmt.Printf("#%-2d: %q\n", i+1, w)
			}
		}

	}

}
