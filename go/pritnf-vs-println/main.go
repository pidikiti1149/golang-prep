package main

import (
	"fmt"
)

func main() {

	//var brand = "Google"

	//var ops, ok, fail int
	//ops, ok, fail = 25, 26, 27

	//fmt.Printf("%s\n", brand)
	//fmt.Printf("%q\n", brand)

	//fmt.Printf("total: %d success: %d/%d\n", ops, ok, fail)

	var speed int
	var heat float64
	var off bool
	var brand string

	fmt.Printf("%T\n%T\n%T\n%T\n", speed, heat, off, brand)

	var (
		planet   = "venus"
		distance = 261
		orbital  = 225.701
		hasLife  = false
	)

	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance is: %v\n", distance)
	fmt.Printf("Orbit range is: %v\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)
	fmt.Printf("%v is %v away. Think! %[2]v kms %[1]v life?\n", planet, distance)
}
