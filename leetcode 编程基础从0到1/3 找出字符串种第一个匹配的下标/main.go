package main

import "fmt"

func main() {
	fmt.Println(buildNext("abcabc"))
}

//	func strStr(haystack string, needle string) int {
//		var j int
//		for i := 0; i+len(needle) <= len(haystack); i++ {
//			j = 0
//			for j < len(needle) {
//				if haystack[i+j] == needle[j] {
//					j++
//					continue
//				}
//			}
//			if j == len(needle) {
//				return i
//			}
//		}
//		return -1
func buildNext(s string) []int {
	n := len(s)
	next := make([]int, n)
	j := 0
	//最匹配前后缀字数
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return next
}
func strStr(haystack string, needle string) int {
	next := buildNext(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != haystack[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}
