package main

import "sync"

/*
利用chan去实现互斥锁，里面存的是空结构体类型
*/
type Mutex struct {
	ch chan struct{}
}

// 这里将我们的chan封装成结构体，方便添加各种方法
// 初始化，实例化一个mutex
func NewMutex() *Mutex {
	mu := &Mutex{ch: make(chan struct{}, 1)}
	return mu
}
func (m *Mutex) Lock() {
	m.ch <- struct{}{}
}
func (m *Mutex) Unlock() {
	<-m.ch
}
func sum(ptrnum *int, wg *sync.WaitGroup, mu *Mutex) {
	defer wg.Done()
	mu.Lock()
	*ptrnum += 1
	mu.Unlock()
}
func main() {
	m := NewMutex()
	var wg sync.WaitGroup
	num := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			sum(&num, &wg, m)
		}()
	}
	wg.Wait()
	println(num)
}
