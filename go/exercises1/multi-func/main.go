package main

import (
	"fmt"
)

func main() {
	a := multi()
	b := demo()

	fmt.Println(a, b)
}

func multi() int {
	return 5
}

func demo() int {
	return 4
}
