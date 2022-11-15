package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("w")
	wg.Wait()

	if msg != "w" {
		t.Error("w tidak ditemukan")
	}
}

func Test_printMessage(t *testing.T) {
	stDout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "w"

	printMessage()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	os.Stdout = stDout
	output := string(result)

	if !strings.Contains(output, "w") {
		t.Error("Cant find w")
	}

}

func Test_main(t *testing.T) {
	stDout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	os.Stdout = stDout
	output := string(result)

	if !strings.Contains(output, "hello farid") {
		t.Error("Cant find farid")
	}
	if !strings.Contains(output, "hello aidil") {
		t.Error("Cant find aidil")
	}
	if !strings.Contains(output, "hello xs") {
		t.Error("Cant find xs")
	}
}
