package main

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int, 0, 2)
	for index, value := range nums {
		if index2, ok := m[target-value]; ok && index != index2 {
			res = append(res, index, index2)
			return res
		}
		m[value] = index
	}
	return nil
}
