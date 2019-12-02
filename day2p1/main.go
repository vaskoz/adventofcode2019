package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// nolint
var (
	in  io.Reader = os.Stdin
	out io.Writer = os.Stdout
)

func main() {
	prog := ""

	fmt.Fscanf(in, "%s", &prog)

	strs := strings.Split(prog, ",")
	opcodes := make([]int, len(strs))

	for i, s := range strs {
		res, _ := strconv.Atoi(s)
		opcodes[i] = res
	}

	opcodes[1] = 12
	opcodes[2] = 2

	for pc := 0; opcodes[pc] != 99; pc += 4 {
		val1 := opcodes[opcodes[pc+1]]
		val2 := opcodes[opcodes[pc+2]]
		outPos := opcodes[pc+3]

		switch code := opcodes[pc]; code {
		case 1:
			opcodes[outPos] = val1 + val2
		case 2:
			opcodes[outPos] = val1 * val2
		}
	}

	fmt.Fprintln(out, opcodes[0])
}
