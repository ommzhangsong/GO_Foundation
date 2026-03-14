package main

import "fmt"

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水
*/
//方法一，用前后缀最大值
func trap(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return 0
	}
	premax := make([]int, n)
	premax[0] = height[0]

	for i := 1; i < n; i++ {
		premax[i] = max(premax[i-1], height[i])
	}
	suf_max := make([]int, n)
	suf_max[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		suf_max[i] = max(suf_max[i+1], height[i])
	}
	for i := 0; i < n-1; i++ {

		ans += min(premax[i], suf_max[i]) - height[i]

	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 方法二，用前后缀最大值 实时
func trap2(height []int) (ans int) {
	left := 0
	n := len(height)
	right := n - 1
	premax := 0
	sufmax := 0
	for left <= right {
		premax = max(height[left], premax)
		sufmax = max(height[right], sufmax)
		if premax <= sufmax {
			ans += premax - height[left]
			left++
			continue
		} else {
			ans += sufmax - height[right]
			right--
			continue
		}
	}
	return ans
}
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap2(height))
}
