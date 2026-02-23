package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var str1 string = "h中国" //7byte,大小16byte（header）
	fmt.Println(len(str1), unsafe.Sizeof(str1))
	str2 := "abcdefeg"

	// 取substring
	substr2 := string(append([]byte(nil), str2[:2]...))
	fmt.Println(substr2)
	// 拼接，高频拼接尽量先转成[]byte再最后转成string

}
