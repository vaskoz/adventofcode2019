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
	goldenOpcodes := make([]int, len(strs))

	for i, s := range strs {
		res, _ := strconv.Atoi(s)
		goldenOpcodes[i] = res
	}

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			opcodes := append([]int{}, goldenOpcodes...)
			opcodes[1] = noun
			opcodes[2] = verb

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

			if opcodes[0] == 19690720 {
				fmt.Fprintln(out, 100*noun+verb)
				return
			}
		}
	}
}
