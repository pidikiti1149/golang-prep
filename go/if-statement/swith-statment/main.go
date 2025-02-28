package main

import (
	"fmt"
)

func main() {

	/*
		city := os.Args[1]

		switch city {
		case "Paris", "Lyon":
			fmt.Println("France")
		case "Toronto":
			fmt.Println("Canada")
		default:
			fmt.Println("Where?")
		}
	*/

	/*i := 0
	switch {

	case i > 0:
		fmt.Println("Positive")
	case i < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("zero")
	}
	*/

	i := 12
	switch {

	case i > 100:
		fmt.Print("Big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number ")
	}

	fmt.Println()
}
