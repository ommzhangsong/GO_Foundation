package main

import "strings"

func repeatedSubstringPattern(s string) bool {
	ss := s + s
	return strings.Contains(ss[1:len(ss)-1], s)

}
