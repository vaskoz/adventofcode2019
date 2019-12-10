package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	in = strings.NewReader("1\n")
	buff := new(bytes.Buffer)
	out = buff

	main()

	result := strings.TrimSpace(buff.String())
	expected := `3512778005`

	if strings.TrimSpace(expected) != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in = strings.NewReader("1\n")
		buff := new(bytes.Buffer)
		out = buff

		main()
	}
}
