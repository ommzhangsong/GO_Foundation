package main

import (
	"context"
	"fmt"
	"time"
)

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%v done", name)
			return
		default:
			fmt.Printf("%v running\n", name)
			time.Sleep(time.Second)
		}

	}
}
func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	// 设置ddl 3s，3s后自动的
	defer cancel()
	go watch(ctx, "G1")
	go watch(ctx, "G2")
	time.Sleep(time.Second * 5)
	fmt.Println("end")
}
