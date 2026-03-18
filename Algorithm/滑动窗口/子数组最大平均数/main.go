package main

import "fmt"

/*
给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。

请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。

任何误差小于 10-5 的答案都将被视为正确答案。
*/
func findMaxAverage(nums []int, k int) (res float64) {
	left := 0

	var sum int
	for _, v := range nums[0:k] {
		sum += v
	}
	res = float64(sum) / float64(k)

	for right := k; right < len(nums); right++ {
		sum += nums[right] - nums[left]
		left++
		res = max(res, float64(sum)/float64(k))
	}
	return res
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
func main() {
	var nums = []int{1, 12, -5, -6, 50, 3}
	fmt.Println(findMaxAverage(nums, 4))
}
