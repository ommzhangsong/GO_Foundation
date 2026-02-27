package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

/*
trace编程，GMP全链路追踪工具，可视化编程，创建启动关闭
*/
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	//逻辑
	fmt.Println("trace.out started")

	//结束
	trace.Stop()
}
