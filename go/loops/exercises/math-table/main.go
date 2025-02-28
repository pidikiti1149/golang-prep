package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	} else if len(args) == 2 {
		fmt.Println("Size is missing")
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}

	m := args[1]
	str1 := "*/+-"

	if strings.IndexAny(str1, m) == -1 {
		fmt.Println("Invalid operator")
		fmt.Println("Valid ops one of: */+-")
		return
	}

	n, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Println("Size is missing")
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	} else if n < 1 {
		fmt.Println("Size is missing")
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}

	fmt.Printf("%5s", m)

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)

	}
	fmt.Println()

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)

		switch m {

		case "+":
			for j := 0; j <= n; j++ {
				fmt.Printf("%5d", i+j)
			}
		case "*":
			for j := 0; j <= n; j++ {
				fmt.Printf("%5d", i*j)
			}
		case "/":
			for j := 0; j <= n; j++ {
				fmt.Printf("%5d", i/j)
			}
		case "-":
			for j := 0; j <= n; j++ {
				fmt.Printf("%5d", i-j)
			}
		}

		fmt.Println()
	}

}
