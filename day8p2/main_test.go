package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	buff := new(bytes.Buffer)
	out = buff

	main()

	result := buff.String()
	expected := `
  f f     f     f   f f f       f f     f       f 
f     f   f     f   f     f   f     f   f       f 
f     f   f     f   f     f   f           f   f   
f f f f   f     f   f f f     f             f     
f     f   f     f   f   f     f     f       f     
f     f     f f     f     f     f f         f     
`

	if expected != result {
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
