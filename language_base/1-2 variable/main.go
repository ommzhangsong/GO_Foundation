package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// there are four ways to state the variables
	//1. var a int , it has default value 0
	//var a int
	//fmt.Println(a)
	////2. we can state and init same time
	//var b int = 100
	//fmt.Println(b)
	////3. 自动识别类型,最简洁,短定义只能用在函数内，不要用全局
	//c := 200
	//fmt.Println(c)
	//fmt.Printf("%T\n", c)
	////4. 自动识别加声明
	//var ab = 1.22
	//fmt.Println(ab)
	//fmt.Printf("%T\n", ab)
	//fmt.Println(int(ab), strconv.Itoa(int(ab)))

	var a string = "abc"
	var b []int = []int{1, 2, 3}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))

	//var b byte = a[0]
	//fmt.Printf("%c", b)

}
