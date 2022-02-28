package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; true; i++ {
		// for j := 0; j < 10; j++ {
		// 	ch <- j * i
		// 	fmt.Printf("producer: %d\n", j*i)
		// }
		ch <- i
		fmt.Printf("producer: %d\n", i)
		time.Sleep(time.Second * 2)
	}
	//close(ch)
}

func consumer(ch <-chan int, str string) {
	// _, isclosed := <-ch
	// fmt.Println(isclosed)
	for i := range ch {
		fmt.Printf("%s %d\n", str, i)
		// time.Sleep(time.Second * 2)
	}
	// for i := 0; true; i++ {
	// 	fmt.Printf("%s %d %d\n", str, i, <-ch)
	// 	time.Sleep(time.Second * 1)
	// }
}

func main() {
	//runtime.GOMAXPROCS(4)
	ch := make(chan int, 10)
	defer close(ch)
	go consumer(ch, "consumer1")
	//go consumer(ch, "consumer2")
	producer(ch)

}
