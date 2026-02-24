package main

import (
	"fmt"
	"time"
)

/*
实现打印123412341234
*/
func main() {
	chs := make([]chan int, 4) // 声明chan数组
	for i := range chs {
		chs[i] = make(chan int) // 4 chan二次初始化
		go func(i int) {
			for {
				v := <-chs[i] // 各读各的
				fmt.Println(v + 1)
				time.Sleep(time.Second)
				chs[(i+1)%4] <- (v + 1) % 4 //下一个赋值谁
			}
		}(i)
	}
	chs[0] <- 0
	select {} //让子线程一直跑
}
