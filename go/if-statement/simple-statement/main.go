package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		n   int
		err error
	)

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number")
	} else if n, err = strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("Cannot convert %q.\n", os.Args[1])
	} else {
		fmt.Printf("%s * 2  %d\n", a[1], n*2)
	}

	fmt.Println("variable n is:", n)
}
