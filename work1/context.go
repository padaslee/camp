package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseCtx := context.Background()

	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)

	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second*5)
	defer cancel()
	go func(c context.Context) {
		ticker := time.NewTicker(time.Second * 1)
		for _ = range ticker.C {
			select {
			case a, ok := <-c.Done():
				fmt.Println("Child exit...", a, ",", ok)
				return
			default:
				fmt.Println("Child running...")
			}
		}
	}(timeoutCtx)

	// range 只读取通道的值，select还会检查通道状态
	// for _ = range timeoutCtx.Done() {
	// 	time.Sleep(time.Second * 1)
	// 	fmt.Println("Main exit...")
	// }
	select {
	case b, ok := <-timeoutCtx.Done():
		time.Sleep(time.Second * 1)
		fmt.Println("Main exit...", b, ",", ok)
	}

	time.Sleep(time.Second * 1)
}
