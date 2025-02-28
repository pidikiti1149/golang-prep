package main

import (
	"fmt"
	"os"
	"strconv"
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

	a := os.Args

	if len(a) != 2 {
		fmt.Println("Give me the magnitude of the earthquake")
		return
	}

	m, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	switch {

	case m < 2.0:
		fmt.Println("micro")
	case m < 3.0:
		fmt.Println("very minor")
	case m < 4.0:
		fmt.Println("minor")
	case m < 5.0:
		fmt.Println("light")
	case m < 6.0:
		fmt.Println("moderate")
	case m < 7.0:
		fmt.Println("strong")
	case m < 8.0:
		fmt.Println("major")
	case m < 10.0:
		fmt.Println("great")
	default:
		fmt.Println("massive")
	}
}
