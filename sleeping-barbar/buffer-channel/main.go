package main

import (
	"fmt"
	"time"
)

func listenTochan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got ", i, " form channel")
		time.Sleep(2 * time.Second)
	}
}
func main() {
	ch := make(chan int, 10)
	go listenTochan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("sending i to chanel")
		ch <- i
		fmt.Println("send ", i, "to channel")
	}
	fmt.Println("Done")
	close(ch)
}
