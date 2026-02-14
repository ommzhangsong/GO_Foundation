package main

// 双指针
func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}
