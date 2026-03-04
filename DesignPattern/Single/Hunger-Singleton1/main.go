package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
懒汉式的创建：这里有俩种，一种是oncedo一种是双重检查锁定
单例模式，一个类只有一个方法，并且这个方法可以被外界调用
*/

type dog struct{}

func (d *dog) bark() { fmt.Println("woowow") }

var lock sync.Mutex
var initialized uint32
var instance *dog

// Getdog 这里用Getdog可以被外部模块所用，并且只暴露一个只读方法，我们千万不能把指针暴露给外部，不然会被修改
func Getdog() *dog {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &dog{}
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}
func main() {
	d := Getdog()
	d.bark()
}
