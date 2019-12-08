package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	wide = 25
	tall = 6
)

// nolint
var (
	out io.Writer = os.Stdout
)

func main() {
	f, _ := os.Open("input.txt")
	image, _ := ioutil.ReadAll(f)
	f.Close()

	final := make([]byte, wide*tall)

	for pos := 0; pos < len(image)-1; {
		for i := pos; i < pos+wide*tall; i++ {
			if final[i-pos] != 0 {
				continue
			}

			switch image[i] {
			case '0':
				final[i-pos] = '0'
			case '1':
				final[i-pos] = '1'
			}
		}

		pos += wide * tall
	}

	fmt.Fprintln(out)

	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			if final[i*wide+j] == '1' {
				fmt.Fprintf(out, "f ")
			} else {
				fmt.Fprintf(out, "  ")
			}
		}
		fmt.Fprintln(out)
	}
}
