package main

import (
	"fmt"
	"math/rand"
)

func main() {

	year := rand.Intn(3000) + 1
	fmt.Println("the year is", year)
	if (year%100 == 0 && year%400 == 0) || (year%100 != 0 && year%4 == 0) {
		fmt.Println("The year is 闰年")
	} else {
		fmt.Println("是平年")
	}

}
