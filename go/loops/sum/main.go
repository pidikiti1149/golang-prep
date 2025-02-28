package main

import (
	"fmt"
)

func main() {

	i := 1
	sum := 0

	for i <= 10 {

		if i%2 != 0 {
			i++
			continue
		}

		i++
		sum += i

	}

	fmt.Println(sum)
}
