package main

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
		if numbers[left]+numbers[right] == target {
			break
		}
	}
	return []int{left, right}
}
