package main

import "fmt"

func printLengthAndCapacity(s []int) {
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}
func main() {
	doAppend := func(s []int) {
		s = append(s, 1)
		printLengthAndCapacity(s)
	}
	s := make([]int, 8)       // s len=cap=8
	doAppend(s[:4])           //长度，容量独立，len = 5, cap = 8
	printLengthAndCapacity(s) // 8,8
	doAppend(s)               // 9,16
	printLengthAndCapacity(s) //9,16
}
