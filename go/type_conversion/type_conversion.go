package main

import "fmt"

func main() {

	speed := 100
	force := 2.5

	//invalid operation: speed * force (mismatched types int and float64)
	//speed = speed * force  (two types cannot be used, need to convert the types)

	speed = speed * int(force) //converting float to int

	fmt.Println(speed)
	fmt.Println(force, int(force))

	{
		speed := 200
		force := 2.5

		speed = int(float64(speed) * force) //converting float to int

		fmt.Println(speed)
	}
}
