package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const hunger = 3

var philosofher = []string{"alif", "aidil", "farid", "zulk", "dewi"}
var wg sync.WaitGroup
var sleeptime = 1 * time.Second
var eatTime = 2 * time.Second
var thinktime = 2 * time.Second
var orderfinish []string
var order sync.Mutex

func dinningProblem(philosofher string, left, right *sync.Mutex) {
	defer wg.Done()

	// Print message
	fmt.Println(philosofher, "is seated")
	time.Sleep(sleeptime)

	for i := hunger; i >= 0; i-- {
		fmt.Println(philosofher, " Is hungry")
		time.Sleep(sleeptime)
		// lock both fork
		left.Lock()
		fmt.Printf("\t%s pickup the fork to his left. \n", philosofher)
		right.Lock()
		fmt.Printf("\t%s pickup the fork to his right. \n", philosofher)

		// print a message
		fmt.Println(philosofher, "has both fork, and is eating .")
		time.Sleep(eatTime)

		// give the pholosofer time to think
		fmt.Println(philosofher, " is thingking")
		time.Sleep(thinktime)

		// unlock and print message
		right.Unlock()
		fmt.Printf("\t%s put down the fork his right.\n", philosofher)
		left.Unlock()
		fmt.Printf("\t%s put down the fork his left. \n", philosofher)

		time.Sleep(sleeptime)
	}
	// print out done message
	fmt.Println(philosofher, "is satisfied")
	time.Sleep(sleeptime)

	fmt.Println(philosofher, "has left the table")
	order.Lock()
	orderfinish = append(orderfinish, philosofher)
	order.Unlock()

}

func main() {
	// print intro
	fmt.Println("the dinning philosofer problem")
	fmt.Println("==============================")
	right := &sync.Mutex{}
	wg.Add(len(philosofher))
	for i := 0; i < len(philosofher); i++ {
		left := &sync.Mutex{}

		go dinningProblem(philosofher[i], left, right)

		right = left
	}

	wg.Wait()
	fmt.Println("the table is empty")
	fmt.Println("==============================")
	fmt.Printf("Order Finished:%s", strings.Join(orderfinish, ","))
}
