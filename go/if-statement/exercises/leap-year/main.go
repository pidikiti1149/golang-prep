package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg_len := len(os.Args[1:])

	if arg_len == 0 {
		fmt.Println("Give me a year number")
		return
	}

	one := os.Args[1]
	year, err := strconv.Atoi(one)

	if err != nil {
		fmt.Println("Invalid year:", os.Args[1])
		return
	}

	if year%4 == 0 && year%100 != 0 {
		fmt.Println(os.Args[1], "is a leap year")
	} else if year%100 == 0 && year%400 == 0 {
		fmt.Println(os.Args[1], "is a leap year")
	} else {
		fmt.Println(os.Args[1], "is not a leap year")
	}
}
