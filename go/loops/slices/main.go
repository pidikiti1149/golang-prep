package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("Hello aditya ela unav nuvu")

	for _, v := range words {

		fmt.Printf("%q\n", v)
	}
}
