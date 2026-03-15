package main

import "fmt"

/*
给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。
*/
func numSubarrayProductLessThanK(nums []int, k int) (count int) {
	if k <= 1 {
		return 0
	}
	l := 0
	r := 0
	res := 1
	for r < len(nums) {
		res = nums[r] * res

		for res >= k {
			res = res / nums[l]
			l++
		}
		count += r - l + 1
		r++

	}
	return count
}
func main() {
	nums := []int{10, 5, 2, 6}
	fmt.Println(numSubarrayProductLessThanK(nums, 100))
}
