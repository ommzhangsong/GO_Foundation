package main

import "fmt"

/*
找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
*/
/*
分析：
长度最小，那我们肯定需要取最优解，定义min
返回的是子数组长度，最小为0，初始化为0,。
我们满足条件是>target,还是一样，只要合法，满足，那我们往右扩展，因为都是解，线性求所有解，我们只求最优解，这里是min
但是，这里我们不需要求所有解，因为我们只需要小的，一旦满足target我们就可以遍历下一轮了，
除此之外，当我们left++，是肯定会下向不合法的路子走，所以是当target>0时，left++
*/
func minSubArrayLen(target int, nums []int) (res int) {
	left, n := 0, len(nums)
	res = len(nums) + 1
	var sum int
	for right := 0; right < n; right++ {
		sum += nums[right]
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}

	}
	if left == 0 {
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
func main() {
	target := 11
	var nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(minSubArrayLen(target, nums))
}
