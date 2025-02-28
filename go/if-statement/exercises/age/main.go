package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Age Seasons
//
//  Let's start simple. Print the expected outputs,
//  depending on the age variable.
//
// EXPECTED OUTPUT
//  If age is greater than 60, print:
//    Getting older
//  If age is greater than 30, print:
//    Getting wiser
//  If age is greater than 20, print:
//    Adulthood
//  If age is greater than 10, print:
//    Young blood
//  Otherwise, print:
//    Booting up
// ---------------------------------------------------------

func main() {

	age := 70

	if age > 60 {
		fmt.Println("Getting Older")
	} else if age > 30 {
		fmt.Println("Getting Wiser")
	} else {
		fmt.Println("Booting UP")
	}
}
