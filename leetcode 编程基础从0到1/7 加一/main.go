package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	// 从后向前加一
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	// 如果最高位也进位，比如 999 -> 1000
	newDigits := make([]int, n+1)
	newDigits[0] = 1
	return newDigits
}

func main() {
	fmt.Println(plusOne([]int{1, 9}))
}
