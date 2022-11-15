package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "hello, world"

	wg.Add(2)
	go updateMessage("hello, univers")
	go updateMessage("hello, bob")

	wg.Wait()
	fmt.Println(msg)
	// maka akan diprint 1 line jika tidak menggunakan mutex
	// go run -race .
	// dan akan menampilkan 1 race artinya 2 go routine mengakses 1 source yang sama
}

// func updateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()
// 	m.Lock()
// 	msg = s
// 	m.Unlock()
// }

// func main() {
// 	msg = "hello, world"

// 	var mu sync.Mutex

// 	wg.Add(2)
// 	go updateMessage("hello, univers", &mu)
// 	go updateMessage("hello, bob", &mu)

// 	wg.Wait()
// 	fmt.Println(msg)
// 	// maka akan diprint 1 line jika tidak menggunakan mutex
// 	// go run -race .
// 	// dan akan menampilkan 1 race artinya 2 go routine mengakses 1 source yang sama
// }
