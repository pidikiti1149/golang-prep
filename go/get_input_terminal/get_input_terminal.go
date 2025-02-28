package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path:", os.Args[0])
	fmt.Println("Second Argument:", os.Args[1])
	fmt.Println("ThirdA rgument:", os.Args[2])
	fmt.Println("Fourth Argument:", os.Args[3])
	fmt.Println("Length of Slice:", len(os.Args))
}
