package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func main() {
	firstName := "Aditya"
	lastName := "Pidikiti"
	temparature := 29.5

	tf := false

	fmt.Printf("My name is %q and my lastname is %s\n", firstName, lastName)
	fmt.Printf("These are %t claims\n", tf)
	fmt.Printf("Temparature is  %g degress\n", temparature)

}
