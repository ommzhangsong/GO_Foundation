package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("1")
		defer fmt.Println("2")
		defer fmt.Println("3")
		return
	}()
	time.Sleep(time.Second * 1)
}
