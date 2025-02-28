package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// var books [4]string
	// var books [1 + 3]string
	data := "2 4 6 1 3 5"

	slc := strings.Split(data, " ")

	var nums []int

	for _, i := range slc {

		num, _ := strconv.Atoi(i)

		nums = append(nums, num)

	}

	fmt.Println("nums        :", nums)

	evens, odds := nums[:3], nums[3:]

	fmt.Println("evens       :", evens)
	fmt.Println("odds        :", odds)

	fmt.Println("middle      :", nums[2:4])
	fmt.Println("first 2     :", nums[:2])
	fmt.Println("last 2      :", nums[len(nums)-2:])

	fmt.Println("evens last 1:", evens[len(evens)-1:])
	fmt.Println("odds last 2 :", odds[len(odds)-2:])

}
