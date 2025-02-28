package main

import (
	"fmt"
)

func main() {
	// var books [4]string
	// var books [1 + 3]string
	var names [3]string
	var distances [5]int
	var data [5]uint8
	var ratios [1]float64
	var alives [4]bool
	var zero [0]uint8

	// print the type
	fmt.Printf("names    : %#v\n", names)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data     : %#v\n", data)
	fmt.Printf("ratios   : %#v\n", ratios)
	fmt.Printf("alives   : %#v\n", alives)
	fmt.Printf("zero     : %#v\n", zero)

	//  2. Print only the types of the same arrays.
	fmt.Println()
	fmt.Printf("names    : %T\n", names)
	fmt.Printf("distances: %T\n", distances)
	fmt.Printf("data     : %T\n", data)
	fmt.Printf("ratios   : %T\n", ratios)
	fmt.Printf("alives   : %T\n", alives)
	fmt.Printf("zero     : %T\n", zero)

	// 3. Print only the elements of the same arrays.
	fmt.Println()
	fmt.Printf("names    : %q\n", names)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("data     : %d\n", data)
	fmt.Printf("ratios   : %.2f\n", ratios)
	fmt.Printf("alives   : %t\n", alives)
	fmt.Printf("zero     : %d\n", zero)
}
