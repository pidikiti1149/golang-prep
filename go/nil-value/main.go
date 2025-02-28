package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	feet := os.Args[1]
	f, err := strconv.Atoi(feet)

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	meters := (float64(f) / 3.281)

	fmt.Printf("%v feet is %.2f meters\n", f, meters)

}
