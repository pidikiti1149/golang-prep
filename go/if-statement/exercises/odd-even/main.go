package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg_len := len(os.Args[1:])

	if arg_len == 0 {
		fmt.Println("Pick a number")
		return
	}

	one := os.Args[1]
	num, err := strconv.Atoi(one)

	if err != nil {
		fmt.Println(os.Args[1], "is not a number")
		return
	}

	if num%2 == 0 && num%8 == 0 {
		fmt.Println(os.Args[1], "is an even number and it's divisible by 8")
	} else if num%2 == 0 {
		fmt.Println(os.Args[1], "is an even number")
	} else if num%2 == 1 {
		fmt.Println(os.Args[1], "is an odd number")
	}
}
