package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	arg_len := len(os.Args[1:])

	if arg_len == 0 {
		fmt.Println("Give me a month name")
		return
	}

	month := strings.ToLower(os.Args[1])
	t := time.Now()
	year := t.Year()

	if month == "january" || month == "march" || month == "may" || month == "july" || month == "august" || month == "october" || month == "december" {
		fmt.Println(os.Args[1], "has 31 days")
	} else if month == "april" || month == "june" || month == "september" || month == "november" {
		fmt.Println(os.Args[1], "has 30 days")
	} else if month == "february" && year%4 == 0 && year%100 != 0 {
		fmt.Println(os.Args[1], "has 29 days")
	} else if month == "february" && year%100 == 0 && year%400 == 0 {
		fmt.Println(os.Args[1], "has 29 days")
	} else if month == "february" {
		fmt.Println(os.Args[1], "has 28 days")
	} else {
		fmt.Println(os.Args[1], "is not a month")
	}

}
