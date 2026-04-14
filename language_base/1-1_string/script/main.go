package main

import (
	"fmt"
)

func main() {
	//
	//var str1 string = "haa" //7byte,大小16byte（header）
	//fmt.Println(len(str1), unsafe.Sizeof(str1))
	str2 := "abcdefeg"

	// 取substring.解决内存滞留的问题
	substr2 := string(append([]byte{}, str2[:2]...))
	fmt.Println(substr2)
	// 拼接，高频拼接尽量先转成[]byte再最后转成string

}
