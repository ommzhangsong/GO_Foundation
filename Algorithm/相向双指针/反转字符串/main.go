package main

import "fmt"

/*
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
*/
func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		str := s[left]
		s[left] = s[right]
		s[right] = str
		left++
		right--
	}
	return
}
func main() {
	str1 := []byte{'a', 'c', 'd', 3, 4}
	reverseString(str1)
	fmt.Println(str1)
	str := []byte("acd") //类型转化。
	reverseString(str)
	fmt.Println(str)
}
