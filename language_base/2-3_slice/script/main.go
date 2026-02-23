package main

import "fmt"

func main() {

	//var brr = []int{1, 2, 3, 4, 5, 6, 7}
	////，append拼接的方法，但会产生中间切片，一种方法去添加元素
	//
	//brr = append(brr[:1], append([]int{2000}, brr[1:]...)...)
	//fmt.Println(brr)
	////	copy方法性能更好
	//var arr = []int{1, 2, 3, 4, 5, 6, 7}
	//arr = append(arr, 0)
	//copy(arr[2:], arr[1:])
	//arr[1] = 2000
	//fmt.Println(arr)
	var a = make([]int, 0, 5)
	a = append(a, 1)
	fmt.Println(len(a), cap(a))

}
