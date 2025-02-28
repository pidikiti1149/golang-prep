// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {

	demo()
}

func demo() {

	s.PrintBacking = true
	s.MaxPerLine = 10

	// #1: assume that tasks can be a long list
	// ---------------------------------------------
	tasks := []string{"jump", "run", "read"}

	upTasks := make([]string, 0, len(tasks))

	s.Show("upTasks", upTasks)

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}
}
