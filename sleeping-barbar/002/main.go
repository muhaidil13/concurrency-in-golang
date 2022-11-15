package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "hello from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "hello from server 2"
	}
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	fmt.Println("Starting server 1 and server 2 ")
	fmt.Println("------------------------------")
	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case s1 := <-channel1:
			fmt.Println("case on s1", s1)
		case s2 := <-channel1:
			fmt.Println("case on s2", s2)
		case s3 := <-channel2:
			fmt.Println("case on s3", s3)
		case s4 := <-channel2:
			fmt.Println("case on s3", s4)

		}
	}
}
