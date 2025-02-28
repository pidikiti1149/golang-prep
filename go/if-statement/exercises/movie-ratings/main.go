package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg_len := len(os.Args[1:])

	if arg_len == 0 {
		fmt.Println("Requires age")
		return
	}

	one := os.Args[1]
	age, err := strconv.Atoi(one)

	if err != nil {
		fmt.Println("Invalid age:", os.Args[1])
		return
	}

	if age > 17 {
		fmt.Println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 && age > 0 {
		fmt.Println("PG-Rated")
	} else if age < 0 {
		fmt.Println(os.Args[1], "Wrong age")
	}
}
