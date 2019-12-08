package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	buff := new(bytes.Buffer)
	out = buff

	main()

	result := strings.TrimSpace(buff.String())
	expected := "2080"

	if strings.TrimSpace(expected) != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buff := new(bytes.Buffer)
		out = buff

		main()
	}
}
