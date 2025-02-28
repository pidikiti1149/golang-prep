package main

import (
	"fmt"
)

func main() {

	/*const (
		monday = iota * 2
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)*/

	/*const (
		EST = -(5 + iota)
		_
		MST
		PST
	)*/

	//fmt.Println(monday, tuesday, wednesday, friday, saturday, sunday)
	//fmt.Println(EST, MST, PST)

	/*const (
		Sep = 9 + iota
		Nov
		Oct
	)*/

	//fmt.Println(Sep, Oct, Nov)

	const (
		_ = iota
		Jan
		Feb
		Mar
	)

	fmt.Println(Jan, Feb, Mar)

	const (
		_ = 3 * iota
		Spring
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
