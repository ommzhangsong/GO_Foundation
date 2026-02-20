package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "time", "值")
	fmt.Println("time:", ctx.Value("time"))
}
