package main

import "fmt"

/*
给你一个整数数组 arr 和 , 整数 k 和 threshold 。

请你返回长度为 k 且平均值大于等于 threshold 的子数组数目。
*/
/*
分析，结果是让我们取返回子数组的个数，所以肯定最小为0，此时初始化为0是没有任何问题的
此外，我们可以看到是计算长度为k的子数组的avg ，就可以用固定长度k的滑动窗口，多一个判断就可以了

*/
func numOfSubarrays(arr []int, k int, threshold int) (res int) {
	left := 0
	right := k
	sum := 0
	for _, v := range arr[0:k] {
		sum += v
	}
	for right < len(arr) {
		if sum/k >= threshold {
			res++
		}
		sum += arr[right] - arr[left]
		left++
		right++
	}
	if sum/k >= threshold {
		res++
	}
	return res
}
func main() {
	var arr = []int{2, 2, 2, 2, 5, 5, 5, 8}
	fmt.Println(numOfSubarrays(arr, 3, 4))
}
