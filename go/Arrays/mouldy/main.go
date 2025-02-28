package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	books := [...][3]string{
		[3]string{"Socrates feels happy", "Socrates feels good", "Socrates feels awesome"},
		[3]string{"Socrates feels sad", "Socrates feels terrible", "Socrates feels bad"},
	}

	args := os.Args

	if len(args) != 3 {
		fmt.Println("Give me two number")
		return
	}

	ver := args[1]
	mood := args[2]

	randomIndex := rand.Intn(len(books[0]))
	books1 := books[0]
	books2 := books[1]
	pos := books1[randomIndex] // get the value from the slice
	neg := books2[randomIndex]

	if ver == "Socrates" && mood == "positive" {
		fmt.Println(pos)
	} else if ver == "Socrates" && mood == "negative" {
		fmt.Println(neg)
	} else {
		fmt.Println("[your name]")
	}

}
