// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

func main() {
	words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
	words = append(words[:1], "is", "everywhere")
	//words = append(words, words[len(words)+1:cap(words)]...)
	fmt.Println("books  :", words[:6])
}
