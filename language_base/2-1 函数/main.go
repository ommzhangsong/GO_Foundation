package main

import (
	"fmt"
)

func trans(arr []string) {
	for i := range arr {
		arr[i] = arr[i] + "new"
	}
}
func main() {
	arr := []string{"a", "b", "c"}
	trans(arr)
	fmt.Println(arr)
}
