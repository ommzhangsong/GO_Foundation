package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Millisecond * 1)
	//创建计时器timer，本身是一个机构体，里面有个time类型的读通道，
	//到达超时时间，就会给C写值,；eg：2026-02-20 14:43:10.1594685 +0800 +08 m=+2.000000001
	//“定时器被标记为到期” ≠ “已经向 timer.C 写完值”,只要调用了 timer.Stop ()，就不要再直接读 timer.C，必须用非阻塞方式。
	// 非阻塞读取，无论超时多久都不会阻塞
	time.Sleep(time.Millisecond * 1)
	select {
	case v := <-timer.C:
		fmt.Printf("成功读取到值：%v\n", v)
	default:
		fmt.Println("通道无值，读取失败（写操作被Stop中断）")
	}
	res := timer.Stop()

	fmt.Printf("%v\n", res)
}
