package main

import "fmt"

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。



示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
*/
/*
分析：首先，是这种子串子数组，我们想到滑动窗口
其次，我们求的是最长子串，所以求最优解大概率需要max
无重复字符，所以我们需要判断这个字符之前是否出现过，判断存在性，那我们要用到map[byte]int
*/
func lengthOfLongestSubstring(s string) (res int) {
	x := []byte(s)
	left, n := 0, len(x)
	if n == 0 {
		return 0
	}
	m := make(map[byte]int)
	for right := 0; right < n; right++ {
		m[x[right]]++
		for m[x[right]] > 1 {
			m[x[left]]--
			left++
		}
		if res < right-left+1 {
			res = right - left + 1
		}
	}
	return res
}
func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
