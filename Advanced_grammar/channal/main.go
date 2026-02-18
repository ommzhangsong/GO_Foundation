package main

import (
	"fmt"
	"time"
)

func main() {

	chan3 := make(chan int, 10)
	chan3 <- 1
	chan3 <- 2
	close(chan3)
	chan3 <- 3
	go func() {
		for range 10 {
			v := <-chan3
			fmt.Println(v)
		}
	}()
	time.Sleep(time.Second * 2)
}
