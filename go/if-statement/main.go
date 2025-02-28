package main

import "fmt"

func main() {

	score := 2
	valid := false

	if score >= 3 && valid {
		fmt.Println("good")
	} else if score < 3 && !valid {
		fmt.Println("bad")
	} else {
		fmt.Println("NoRating")
	}
}
