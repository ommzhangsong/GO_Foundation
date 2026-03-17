package main

import "fmt"

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
*/
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func lengthOfLongestSubstring(s string) (res int) {
	m := make(map[byte]int)
	x := []byte(s)
	if len(x) == 0 {
		return 0
	}
	if len(x) == 1 {
		return 1
	}
	left := 0
	right := 0
	res = 0
	for right < len(x) {
		m[x[right]]++
		for m[x[right]] > 1 {
			m[x[left]]--
			left++
		}
		res = max(res, right-left+1)
		right++
	}
	return res
}
func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
