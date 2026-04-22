package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	// 1. sort库，传统，自定义匿名函数，一般降序时使用
	slice1 := []int{23, 31, 4, 22, 42, 12, 43, 23, 54, 65, 14}
	sort.Slice(slice1, func(i, j int) bool { return slice1[i] > slice1[j] })
	fmt.Println(slice1)
	// slices库常用。o(nlogn)
	slice1 = []int{23, 31, 4, 22, 42, 12, 43, 23, 54, 65, 14}
	slices.Sort(slice1)
	fmt.Println(slice1)
}
