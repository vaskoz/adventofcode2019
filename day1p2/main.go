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
		for fuel := (mass / 3) - 2; fuel > 0; fuel = (fuel / 3) - 2 {
			totalFuel += fuel
		}
	}
	fmt.Fprintln(out, totalFuel)
}
