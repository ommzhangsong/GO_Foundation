package main

import "fmt"

func main() {
	result := findTheDifference("abv", "xbav")
	fmt.Println(result)
}

// 方法一，用减法，最后用int->byte类型打印出字母
//
//	func findTheDifference(s string, t string) byte {
//		ssum, tsum := 0, 0
//		for _, x := range s {
//			ssum += int(x)
//		}
//		for _, x := range t {
//			tsum += int(x)
//		}
//		return byte(tsum - ssum)
//	}
//
// 第二种 map计数法
// func findTheDifference(s string, t string) byte {
//
//		bucket := make(map[byte]int)
//		for i := 0; i < len(s); i++ {
//			bucket[s[i]]++
//		}
//		for i := 0; i < len(t); i++ {
//			bucket[t[i]]--
//			if bucket[t[i]] < 0 {
//				return t[i]
//			}
//		}
//		return byte(0)
//	}
//
// 异或方法
func findTheDifference(s string, t string) byte {
	var res byte
	for i := 0; i < len(s); i++ {
		res ^= s[i]
		res ^= t[i]
	}
	res ^= t[len(t)-1]
	return res
}
