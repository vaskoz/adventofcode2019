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
		doubleSeen := false
		valid := true
		candidate := i
		lastDigit := candidate % 10
		candidate /= 10

		for candidate != 0 {
			d := candidate % 10
			if d == lastDigit {
				doubleSeen = true
			} else if d > lastDigit {
				valid = false
				break
			}

			lastDigit = d
			candidate /= 10
		}

		if doubleSeen && valid {
			validPasswords++
		}
	}

	fmt.Fprintln(out, validPasswords)
}
