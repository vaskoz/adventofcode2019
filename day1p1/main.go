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
	var mass, totalFuel int

	for _, err := fmt.Fscanf(in, "%d", &mass); err == nil; _, err = fmt.Fscanf(in, "%d", &mass) {
		totalFuel += (mass / 3) - 2
	}
	fmt.Fprintln(out, totalFuel)
}
