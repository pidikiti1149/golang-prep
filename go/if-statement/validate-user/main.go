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
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage : [username] [password]")
		return
	}

	const (
		AccessGranted    = "Access granted for %q\n"
		AccessDenied     = "Access denied for %q\n"
		InvalidPassword  = "Invalid password %q\n"
		wrongCredentials = "Entered wrong credentials\n"

		username1 = "jack"
		username2 = "tom"

		password1 = 1888
		password2 = 1997
	)
	username := os.Args[1]
	pass := os.Args[2]

	password, err := strconv.Atoi(pass)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	if username == username1 && password != password1 {
		fmt.Printf(InvalidPassword, username)
	} else if username == username2 && password != password2 {
		fmt.Printf(InvalidPassword, username)
	} else if username == username1 && password == password1 {
		fmt.Printf(AccessGranted, username)
	} else if username == username2 && password == password2 {
		fmt.Printf(AccessGranted, username)
	} else {
		fmt.Printf(wrongCredentials)
	}

}
