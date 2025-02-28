// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

const (
	x, y, z = "lower", "upper", "title"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("usage")
		return
	}

	inp1, inp2 := args[1], args[2]

	switch {

	case x == inp1:
		fmt.Println(strings.ToLower(inp2))
	case y == inp1:
		fmt.Println(strings.ToUpper(inp2))
	case z == inp1:
		fmt.Println(strings.ToTitle(inp2))
	default:
		fmt.Printf("Unknown command: %q\n", inp1)
	}
}
