package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 1)
	select {
	case <-timer1.C:
		fmt.Println("Time out...")
	}

	timer2 := time.NewTicker(time.Second * 1)
	for i := 0; i < 10; i++ {
		select {
		case <-timer2.C:
			fmt.Println("Ticker...")
		}
	}
	for _ = range timer2.C {
		fmt.Println("Ticker2...")
	}
}
