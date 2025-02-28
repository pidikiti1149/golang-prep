package main

import (
	"fmt"
)

func main() {
	// var books [4]string
	// var books [1 + 3]string
	var books []string

	books = []string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	// print the type
	fmt.Printf("books  : %T", books)

	fmt.Printf(" %d", len(books))
	fmt.Printf(" %t", books == nil)

	fmt.Println()

}
