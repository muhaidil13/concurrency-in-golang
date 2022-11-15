package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(m string) {
	defer wg.Done()
	msg = m
}

func printMessage() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {
	msg = "hello world"
	wg.Add(1)
	go updateMessage("hello farid")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("hello aidil")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("hello xs")
	wg.Wait()
	printMessage()
}
