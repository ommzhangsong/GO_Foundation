package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(100) + 1
	fmt.Println("enter a number")
	var a int
	for {
		fmt.Scan(&a)
		if a > num {
			fmt.Println("You are too big")

		} else if a < num {
			fmt.Println("You are too small")

		} else {
			fmt.Println("right!!")
			break
		}
	}
}
