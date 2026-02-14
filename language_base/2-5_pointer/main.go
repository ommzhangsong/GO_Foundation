package main

import "fmt"
import "unsafe"

func main() {
	var a *int
	var b int = 100
	a = &b
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(a)

}
