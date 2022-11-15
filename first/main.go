package main

import (
	"fmt"
	"sync"
)

func printSometing(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
func main() {
	// using waitgroup
	var wg sync.WaitGroup

	// bad way if using time.sleep
	// go printSometing("Ini yang akan diprint pertama")

	word := []string{
		"s",
		"sa",
		"sdsd",
		"asd1",
		"sdsds11",
		"aee22",
		"sd2232",
		"jhjhj",
		"tytyty",
	}
	wg.Add(len(word))
	for _, val := range word {
		go printSometing(val, &wg)
	}

	wg.Wait()

	// time.Sleep(3 * time.Second)

	wg.Add(1)
	printSometing("Ini yang akan diprint Kedua", &wg)
}
