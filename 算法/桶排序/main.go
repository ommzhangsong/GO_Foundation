package main

import "fmt"

func main() {
	//桶排序算法，一个数字一个桶的方法，也叫计数排序,时间换空间
	arr := []int{11, 22, 23, 43, 5, 65, 3, 67, 59, 99}
	max := arr[0]
	//想找到最大的，最大99那么我们就需要100个桶，0-99
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	count := make([]int, max+1)
	//为什么计数，怕重复
	for _, v := range arr {
		count[v]++
	}
	//再次创建一个切片，把位置有数字直接加上，这里的数字就是索引i
	sorted := make([]int, 0, len(arr))
	for i, _ := range count {
		for count[i] > 0 {
			sorted = append(sorted, i)
			count[i]--
		}
	}
	fmt.Println(sorted)
}
