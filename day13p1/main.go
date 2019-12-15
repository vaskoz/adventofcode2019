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
	sb := &strings.Builder{}
	runProgram(sb)

	lines := strings.Split(sb.String(), "\n")
	lines = lines[:len(lines)-1]

	count := 0

	for pos := 0; pos < len(lines); pos += 3 {
		if lines[pos+2] == "2" {
			count++
		}
	}

	fmt.Fprintln(out, count)
}

func runProgram(out io.Writer) {
	f, _ := os.Open("input.txt")
	defer f.Close()
	data, _ := ioutil.ReadAll(f)
	prog := string(data)
	prog = strings.TrimSpace(prog)

	strs := strings.Split(prog, ",")
	opcodes := make([]int, 1000*len(strs))

	for i, s := range strs {
		res, _ := strconv.Atoi(s)
		opcodes[i] = res
	}

	relativeBase := 0

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
		case 5: // jump-if-true
			if opcodes[opcodes[pc+1]] != 0 {
				pc = opcodes[opcodes[pc+2]]
			} else {
				pc += 3
			}
		case 6: // jump-if-false
			if opcodes[opcodes[pc+1]] == 0 {
				pc = opcodes[opcodes[pc+2]]
			} else {
				pc += 3
			}
		case 7: // less-than
			if opcodes[opcodes[pc+1]] < opcodes[opcodes[pc+2]] {
				opcodes[opcodes[pc+3]] = 1
			} else {
				opcodes[opcodes[pc+3]] = 0
			}

			pc += 4
		case 8: // equals
			if opcodes[opcodes[pc+1]] == opcodes[opcodes[pc+2]] {
				opcodes[opcodes[pc+3]] = 1
			} else {
				opcodes[opcodes[pc+3]] = 0
			}

			pc += 4
		case 9: // update relative base
			relativeBase += opcodes[opcodes[pc+1]]

			pc += 2
		default: // deal with parameter modes
			var val1, val2 int

			op := code % 100
			code /= 100

			switch op {
			case 1, 2:
				if code%10 == 2 {
					val1 = opcodes[relativeBase+opcodes[pc+1]]
				} else if code%10 == 1 {
					val1 = opcodes[pc+1]
				} else {
					val1 = opcodes[opcodes[pc+1]]
				}

				code /= 10
				if code%10 == 2 {
					val2 = opcodes[relativeBase+opcodes[pc+2]]
				} else if code%10 == 1 {
					val2 = opcodes[pc+2]
				} else {
					val2 = opcodes[opcodes[pc+2]]
				}
				code /= 10

				var outPos int

				if code%10 == 2 {
					outPos = relativeBase + opcodes[pc+3]
				} else {
					outPos = opcodes[pc+3]
				}

				switch op {
				case 1:
					opcodes[outPos] = val1 + val2
				case 2:
					opcodes[outPos] = val1 * val2
				}

				pc += 4
			case 3:
				var userInput, writeTo int

				fmt.Fscanf(in, "%d", &userInput)

				if code == 1 {
					writeTo = opcodes[pc+1]
				} else if code == 2 {
					writeTo = relativeBase + opcodes[pc+1]
				}
				opcodes[writeTo] = userInput
				pc += 2

			case 4:
				if code == 1 {
					fmt.Fprintln(out, opcodes[pc+1])
				} else if code == 2 {
					fmt.Fprintln(out, opcodes[relativeBase+opcodes[pc+1]])
				}

				pc += 2
			case 5: // jump-if-true
				var v1, v2 int
				if code%10 == 2 {
					v1 = opcodes[relativeBase+opcodes[pc+1]]
				} else if code%10 == 1 {
					v1 = opcodes[pc+1]
				} else {
					v1 = opcodes[opcodes[pc+1]]
				}

				code /= 10

				if code%10 == 2 {
					v2 = opcodes[relativeBase+opcodes[pc+2]]
				} else if code%10 == 1 {
					v2 = opcodes[pc+2]
				} else {
					v2 = opcodes[opcodes[pc+2]]
				}

				if v1 != 0 {
					pc = v2
				} else {
					pc += 3
				}
			case 6: // jump-if-false
				var v1, v2 int
				if code%10 == 2 {
					v1 = opcodes[relativeBase+opcodes[pc+1]]
				} else if code%10 == 1 {
					v1 = opcodes[pc+1]
				} else {
					v1 = opcodes[opcodes[pc+1]]
				}

				code /= 10
				if code%10 == 2 {
					v2 = opcodes[relativeBase+opcodes[pc+2]]
				} else if code%10 == 1 {
					v2 = opcodes[pc+2]
				} else {
					v2 = opcodes[opcodes[pc+2]]
				}

				if v1 == 0 {
					pc = v2
				} else {
					pc += 3
				}
			case 7: // less-than
				var v1, v2 int

				if code%10 == 2 {
					v1 = opcodes[relativeBase+opcodes[pc+1]]
				} else if code%10 == 1 {
					v1 = opcodes[pc+1]
				} else {
					v1 = opcodes[opcodes[pc+1]]
				}

				code /= 10

				if code%10 == 2 {
					v2 = opcodes[relativeBase+opcodes[pc+2]]
				} else if code%10 == 1 {
					v2 = opcodes[pc+2]
				} else {
					v2 = opcodes[opcodes[pc+2]]
				}

				code /= 10

				var outPos int

				if code%10 == 2 {
					outPos = relativeBase + opcodes[pc+3]
				} else {
					outPos = opcodes[pc+3]
				}

				if v1 < v2 {
					opcodes[outPos] = 1
				} else {
					opcodes[outPos] = 0
				}

				pc += 4
			case 8: // equals
				var v1, v2 int

				if code%10 == 2 {
					v1 = opcodes[relativeBase+opcodes[pc+1]]
				} else if code%10 == 1 {
					v1 = opcodes[pc+1]
				} else {
					v1 = opcodes[opcodes[pc+1]]
				}

				code /= 10

				if code%10 == 2 {
					v2 = opcodes[relativeBase+opcodes[pc+2]]
				} else if code%10 == 1 {
					v2 = opcodes[pc+2]
				} else {
					v2 = opcodes[opcodes[pc+2]]
				}
				code /= 10

				var outPos int

				if code%10 == 2 {
					outPos = relativeBase + opcodes[pc+3]
				} else {
					outPos = opcodes[pc+3]
				}

				if v1 == v2 {
					opcodes[outPos] = 1
				} else {
					opcodes[outPos] = 0
				}

				pc += 4
			case 9:
				if code%10 == 2 {
					relativeBase += opcodes[relativeBase+opcodes[pc+1]]
				} else if code%10 == 1 {
					relativeBase += opcodes[pc+1]
				} else {
					relativeBase += opcodes[opcodes[pc+1]]
				}

				pc += 2
			}
		}
	}
}
