package main

import (
	"fmt"
	"strings"
)

func main() {
	// var books [4]string
	// var books [1 + 3]string
	names := [...]string{"Einstein", "Tesla", "Shepard"}
	distances := [...]int{30, 40, 70, 77, 90}

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)
	for i, n := range names {
		fmt.Printf("names[%d]  : %#v\n", i, n)
	}

	fmt.Print("\ndistances", separator)
	for i, n := range distances {
		fmt.Printf("distances[%d]  : %#v\n", i, n)
	}
}
