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

	var ans, pos int

	fewestZeros := int(^uint(0) >> 1)

	for pos < len(image)-1 {
		var zeros, ones, twos int

		for i := pos; i < pos+wide*tall; i++ {
			switch image[i] {
			case '0':
				zeros++
			case '1':
				ones++
			case '2':
				twos++
			}
		}

		pos += wide * tall

		if zeros < fewestZeros {
			fewestZeros = zeros
			ans = ones * twos
		}
	}

	fmt.Fprintln(out, ans)
}
