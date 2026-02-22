package main

import (
	"fmt"
	"sync"
)

func main() {

	chan3 := make(chan int, 10) // 有缓冲chan
	var wg sync.WaitGroup
	//type wchannel chan<- int //只能写
	//type rchannel <-chan int //只能读，且不能关闭
	//fmt.Println(wchannel(chan3), rchannel(chan3))
	chan3 <- 1
	chan3 <- 2
	chan3 <- 3
	wg.Add(1)
	close(chan3)
	//关闭，关闭语义只需关闭一次就好了，发生panic提醒程序员逻辑错误
	//关闭后继续写，也是panic
	//关闭后是可以继续读的，读的是零值。且永远不会阻塞，卡死永远循环，用ok语法，或者range chan
	go func() {
		defer wg.Done()
		for v := range chan3 {

			fmt.Println(v)
		}
	}()
	wg.Wait()
}
