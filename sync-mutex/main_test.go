package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	a, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(a)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "530400.00") {
		t.Error("wrong values")
	}
}
