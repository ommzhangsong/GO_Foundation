package main

import (
	"fmt"
	"sync"
)

type Resource struct {
	sync.RWMutex
	m map[int]int
}

func main() {

	/*
		关于map底层，map本质是指针，指向hmap，底层是以桶算法实现的哈希表
		当装载因子过高，或者overflow桶过多，会发生扩容。
		扩容：2倍扩容，原地扩容
		2倍扩容会导致map地址发生变化，所以如果同时读写，肯定会发生一致性问题，
		也就是读到了map的旧地址（为空），所以这里加上RW mutex锁
	*/
	// 读函数
	var wg sync.WaitGroup
	wg.Add(2)
	r := Resource{m: make(map[int]int)}
	go func() {
		defer wg.Done()
		for j := range 100 {
			r.RLock()
			fmt.Println(r.m[j])
			r.RUnlock()
		}
	}()
	go func() {
		defer wg.Done()
		for j := range 100 {
			r.Lock()
			r.m[j] = j
			r.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println("done")
}
