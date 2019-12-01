package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	file, _ := os.Open("input.txt")
	in = file
	buff := new(bytes.Buffer)
	out = buff

	main()

	result := strings.TrimSpace(buff.String())
	if expected := "5049684"; expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, _ := os.Open("input.txt")
		in = file
		buff := new(bytes.Buffer)
		out = buff

		main()
	}
}
