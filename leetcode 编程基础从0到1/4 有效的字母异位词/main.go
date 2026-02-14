package main

import (
	"fmt"
	"sort"
)

// map计算法
//
//	func isAnagram(s string, t string) bool {
//		bucket := make(map[byte]int)
//		if len(s) != len(t) {
//			return false
//		}
//		for i := 0; i < len(s); i++ {
//			bucket[s[i]]++
//		}
//		for j := 0; j < len(t); j++ {
//			bucket[t[j]]--
//			if bucket[t[j]] < 0 {
//				return false
//			}
//		}
//		return true
//	}
//
// 排序后再比较
func isAnagram(s string, t string) bool {
	s1 := []byte(s)
	s2 := []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s2[j] })
	sort.Slice(s2, func(i, j int) bool { return s1[i] < s2[j] })
	return string(s1) == string(s2)

}
func main() {
	fmt.Println(isAnagram("abcd", "adbc"))
}
