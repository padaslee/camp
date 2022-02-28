package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Send to channel...")
		ch <- 1
	}()

	<-ch
	fmt.Println("Receiev from channel...")
}
