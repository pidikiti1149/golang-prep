package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Fix
//
//  1. Uncomment the code
//
//  2. And fix the problems
//
//  3. BONUS: Simplify the code
// ---------------------------------------------------------

func main() {

	students := [2][3]float64{
		[3]float64{3, 5, 6},
		[3]float64{6, 7, 8},
	}

	var sum float64
	sum += students[0][0] + students[0][1] + students[0][2]
	sum += students[1][0] + students[1][1] + students[1][2]

	const N = float64(len(students) * len(students[0]))

	fmt.Println("Avg grade:", sum/N)

}
