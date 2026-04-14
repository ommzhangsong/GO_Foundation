package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	//1. ❗并不是必须项，它是权衡，一般我们会选择性能，所以做索引切片取subslice：copy去做subsilce，因为可以省内存，不共享内存，不会存在滞留内存，强行做内存复制，所以这是一个性能与内存的权衡
	sub := make([]int, 3)
	copy(sub, arr[0:3])
	fmt.Println("这是subslice:", sub)
	// 2. copy去做删除头尾以外，特定索引下标的元素，
	fmt.Println("删除最后一个元素后：", arr[:len(arr)-1])
	copy(arr[2:], arr[3:])
	fmt.Println("删特定下标一个元素后：", arr[:len(arr)-1])
	//3. copy去做整体的移动,特别是插入，腾出来一个位置
	arr = []int{1, 2, 3, 4, 5, 6}
	copy(arr[3:], arr[2:])
	fmt.Println("腾出来一个位置，腾出来索引为2的位置,就可以在2插入了", arr)
}
