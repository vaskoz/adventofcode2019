package main

import (
	"fmt"
	"io"
	"os"
)

// nolint
var (
	in  io.Reader = os.Stdin
	out io.Writer = os.Stdout
)

func main() {
	validPasswords := 0

	for i := 372037; i <= 905157; i++ {
		valid := true
		candidate := i
		lastDigit := candidate % 10
		counts := make(map[int]int)
		counts[lastDigit]++

		candidate /= 10

		for candidate != 0 {
			d := candidate % 10
			if d > lastDigit {
				valid = false
				break
			}

			counts[d]++
			lastDigit = d
			candidate /= 10
		}

		if valid {
			for _, freq := range counts {
				if freq == 2 {
					validPasswords++
					break
				}
			}
		}
	}

	fmt.Fprintln(out, validPasswords)
}
