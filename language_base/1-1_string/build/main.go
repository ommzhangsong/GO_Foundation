package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	   	它的“黑科技”：
	      零拷贝转换：最后当你调用 b.String() 时，它通过 unsafe.Pointer 直接把底层的 []byte 强转成 string 返回。它不需要像 bytes.Buffer 那样重新申请一次内存去拷贝数据。

	      只增不减：它只支持写入，不支持修改中间的数据，这保证了其逻辑的简洁和高效。
	*/
	var b strings.Builder
	fmt.Printf("初始: len=%d, cap=%d\n", b.Len(), b.Cap())

	b.Grow(10) // 预分配 10 字节，来支持预定义容量。当我们可以预定义我们需要使用的容量时，strings.Builder 就能避免扩容而创建新的 slice 了。
	fmt.Printf("Grow后: len=%d, cap=%d\n", b.Len(), b.Cap())
	//如果已经有容量长度再grow，通过 current_capacity * 2 + n （n 就是你想要扩充的容量）的方式来对内部的 slice 进行扩容的。
	b.WriteString("hello")
	a := b
	fmt.Println(a)
	fmt.Printf("写入后: len=%d, cap=%d\n", b.Len(), b.Cap())

	// 再次写入超过容量的内容
	b.WriteString(" world")
	fmt.Printf("溢出后: len=%d, cap=%d\n", b.Len(), b.Cap())
	fmt.Fprint(&b, "得", 100, true) //拼接各种类型，直接输入，不过效率比writing要低
	fmt.Println(b.String())
}
