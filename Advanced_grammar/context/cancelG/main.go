package main

import (
	"context"
	"fmt"
	"time"
)

/*
用来取消携程
*/
func watch(ctx context.Context, name string) { // context包下的Context接口类型，
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Printf("%v running\n", name)
			time.Sleep(time.Second)
		}

	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "goroutine1")
	go watch(ctx, "goroutine2")
	time.Sleep(time.Second * 2)
	cancel() //一个cancel同时可以控制俩个goroutine，select监听chan状态
	time.Sleep(time.Second * 1)

}
