package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
	f, _ := os.Open("input.txt")
	defer f.Close()
	data, _ := ioutil.ReadAll(f)
	prog := string(data)
	prog = strings.TrimSpace(prog)

	strs := strings.Split(prog, ",")
	opcodes := make([]int, len(strs))

	for i, s := range strs {
		res, _ := strconv.Atoi(s)
		opcodes[i] = res
	}

	for pc := 0; opcodes[pc] != 99; {
		switch code := opcodes[pc]; code {
		case 1:
			val1 := opcodes[opcodes[pc+1]]
			val2 := opcodes[opcodes[pc+2]]
			outPos := opcodes[pc+3]
			opcodes[outPos] = val1 + val2
			pc += 4
		case 2:
			val1 := opcodes[opcodes[pc+1]]
			val2 := opcodes[opcodes[pc+2]]
			outPos := opcodes[pc+3]
			opcodes[outPos] = val1 * val2
			pc += 4
		case 3:
			var userInput int

			fmt.Fscanf(in, "%d", &userInput)

			writeTo := opcodes[pc+1]
			opcodes[writeTo] = userInput
			pc += 2
		case 4:
			readFrom := opcodes[pc+1]
			fmt.Fprintln(out, opcodes[readFrom])

			pc += 2
		default: // deal with parameter modes
			var val1, val2 int

			op := code % 100
			code /= 100

			switch op {
			case 1, 2:
				if code%10 == 1 {
					val1 = opcodes[pc+1]
				} else {
					val1 = opcodes[opcodes[pc+1]]
				}

				code /= 10
				if code%10 == 1 {
					val2 = opcodes[pc+2]
				} else {
					val2 = opcodes[opcodes[pc+2]]
				}

				outPos := opcodes[pc+3]

				switch op {
				case 1:
					opcodes[outPos] = val1 + val2
				case 2:
					opcodes[outPos] = val1 * val2
				}

				pc += 4
			case 4:
				if code == 1 {
					fmt.Fprintln(out, opcodes[pc+1])
				}

				pc += 2
			}
		}
	}
}
