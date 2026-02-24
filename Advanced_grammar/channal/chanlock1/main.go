package main

import "sync"

/*
主要依靠有缓冲锁容量为1，因为只要没读走，下一个读会严格阻塞
这样我们可以实现锁的功能，
*/
func sum(ch chan struct{}, num *int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- struct{}{}
	*num++
	<-ch
}
func main() {
	ch := make(chan struct{}, 1)
	var wg sync.WaitGroup
	wg.Add(100)
	num := 0
	for i := 0; i < 100; i++ {
		go sum(ch, &num, &wg)
	}

	wg.Wait()
	println(num)
}
