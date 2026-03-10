package main

import (
	"fmt"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a ，b ，c ，使得 a + b + c = 0 ？请找出所有和为 0 且 不重复 的三元组。


*/
/*
思路：我们可以看到让我们返回的是
 value而不是index所以我们是可以进行排序的
为什么要排序是我们想用双向的指针，就像两数只和那样。*/
func threeSum(nums []int) [][]int {

	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {

		x := nums[i]
		if i > 0 && x == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			s := x + nums[left] + nums[right]
			if s > 0 {
				right--
			} else if s < 0 {
				left++
			} else {
				res = append(res, []int{x, nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}

	}
	return res
}

func main() {
	var a = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(a))

}
