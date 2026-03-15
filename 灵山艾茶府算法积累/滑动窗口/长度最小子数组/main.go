package main

/*
找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
*/
func minSubArrayLen(target int, nums []int) (res int) {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left := 0
	res = n + 1
	sum := 0
	right := 0
	for right < n {
		sum += nums[right]
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if res == n+1 {
		return 0
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
