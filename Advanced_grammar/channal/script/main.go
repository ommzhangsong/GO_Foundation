package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 20
	go func() {
		fmt.Println("携程先执行")
		<-ch
	}()

}
