package main

import "strconv"

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool, len(s))
	maxlen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		for m[s[right]] {
			left++
			delete(m, m[left])
		}
	}

}
