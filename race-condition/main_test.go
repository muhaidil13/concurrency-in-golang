package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "hello world"
	wg.Add(2)
	go updateMessage("goodbye world")
	go updateMessage("goodbye worl")
	wg.Wait()

	if msg != "goodbye world" {
		t.Error("cant find test")
	}
}
