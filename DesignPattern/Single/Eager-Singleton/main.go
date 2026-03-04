package main

import "fmt"

/*
饿汉式的创建
单例模式，一个类只有一个方法，并且这个方法可以被外界调用
*/

type dog struct{}

func (d *dog) bark() { fmt.Println("woowow") }

var ptrdog = &dog{}

// Getdog 这里用Getdog可以被外部模块所用，并且只暴露一个只读方法，我们千万不能把指针暴露给外部，不然会被修改
func Getdog() *dog {
	return ptrdog
}
func main() {
	d := Getdog()
	d.bark()
}
