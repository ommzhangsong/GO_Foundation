package main

import (
	"fmt"
	"sync"
)

/*
chan 实现的限流器，比如不管在什么时候我只有3个goroutine在执行
*/

func main() {
	ch := make(chan struct{}, 3)
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		ch <- struct{}{} //第三个就阻塞了，就不会再开线程了，就只开了三个线程
		// 重点是他不是用chan去做什么存储数据，chan存数据不重要，
		//❗使用chan的阻塞机制，让每次只能开3个线程，每次处理3个逻辑事务
		go func(i int) {
			defer wg.Done()
			fmt.Printf("正在执行第%d的业务\n ", i)
			<-ch
		}(i)
	}
	wg.Wait()
	fmt.Println("done")
}
