package main

import (
	"fmt"
	"strings"
)

func main() {

	scientists := [2][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
	}

	separator := "\n" + strings.Repeat("=", 40) + "\n"

	// Header row
	fmt.Printf("%-15s %-15s %-10s\n", "First Name", "Last Name", "Nickname")
	fmt.Println(separator)

	// Printing data in formatted columns
	for _, s := range scientists {
		fmt.Printf("%-15s %-15s %-10s\n", s[0], s[1], s[2])
	}
}
