package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	sleeptime = 0 * time.Second
	eatTime = 0 * time.Second
	thinktime = 0 * time.Second

	for i := 0; i <= 100; i++ {
		main()
		if len(orderfinish) != 5 {
			t.Error("error wrong number")
		}
		orderfinish = []string{}
	}

}
