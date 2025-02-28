package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("Give me correct numbers")
		return
	}

	min, err := strconv.Atoi(args[1])

	max, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Println("Give only two integers")
		return
	}

	sum := 0
	i := min

	for {

		if i%2 != 0 {
			i++
			continue
		}

		if i > max {
			break
		}

		sum += i

		fmt.Print(i)

		if i < max {
			fmt.Print(" + ")
		}

		i++

	}

	fmt.Printf(" = %d\n", sum)
}
