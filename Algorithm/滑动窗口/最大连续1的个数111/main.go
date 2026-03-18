package main

import "fmt"

/*
给定一个二进制数组 nums 和一个整数 k，假设最多可以翻转 k 个 0 ，则返回执行操作后 数组中连续 1 的最大个数 。
*/
/*
返回最大连续1的个数,最小ans是0，所以可以初始化为0
边界判断，无
可变的滑动窗口，right扩大窗口去找1，left减小窗口，去减少0，俩者可以同时从0开始找
*/
func longestOnes(nums []int, k int) (ans int) {
	left := 0
	zero := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zero++
		}
		for zero > k {
			if nums[left] == 0 {
				zero--
			}
			left++

		}
		ans = max(ans, right-left+1)
	}

	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	var nums = []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	fmt.Println(longestOnes(nums, k))
}
