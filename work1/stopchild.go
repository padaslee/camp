package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)
	go func() {
		select {
		case <-done:
			fmt.Println("Child exit...")
			return
		}
	}()
	fmt.Println("Main exit...")
	close(done)
	time.Sleep(time.Second * 1)
}
