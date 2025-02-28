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
	t := 50

	switch {

	case t > 4:
		fmt.Println("good morning")
	case t > 12:
		fmt.Println("good afternoon")
	case t > 16:
		fmt.Println("good evening")
	case t > 20:
		fmt.Println("good night")
	default:
		fmt.Println("good night")
	}
}
