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
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {

	one := os.Args[1]
	con := strings.ContainsAny(one, "bcdfghjklmnpqrstvwz")
	vow := strings.ContainsAny(one, "aeiou")
	sem := strings.ContainsAny(one, "xy")
	if len(os.Args[1]) == 1 && vow {
		fmt.Printf("%q is a vowel\n", one)
	} else if len(os.Args[1]) == 1 && con {
		fmt.Printf("%q is a consonant\n", one)
	} else if len(os.Args[1]) == 1 && sem {
		fmt.Printf("%q is sometimes a vowel, sometimes not\n", one)
	} else {
		fmt.Println("Give me a letter")
	}

}
