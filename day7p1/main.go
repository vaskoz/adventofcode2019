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
	used := make(map[int]struct{})
	var maxThurst int

	for a := 0; a < 5; a++ {
		used[a] = struct{}{}
		for b := 0; b < 5; b++ {
			if _, use := used[b]; use {
				continue
			}
			used[b] = struct{}{}
			for c := 0; c < 5; c++ {
				if _, use := used[c]; use {
					continue
				}
				used[c] = struct{}{}
				for d := 0; d < 5; d++ {
					if _, use := used[d]; use {
						continue
					}
					used[d] = struct{}{}
					for e := 0; e < 5; e++ {
						if _, use := used[e]; use {
							continue
						}
						thrust := 0
						thrust = run([]int{a, 0})
						thrust = run([]int{b, thrust})
						thrust = run([]int{c, thrust})
						thrust = run([]int{d, thrust})
						thrust = run([]int{e, thrust})

						if thrust > maxThurst {
							maxThurst = thrust
						}
					}
					delete(used, d)
				}
				delete(used, c)
			}
			delete(used, b)
		}
		delete(used, a)
	}
	fmt.Fprintln(out, maxThurst)
}

func run(in []int) int {
	inPos := 0
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
			writeTo := opcodes[pc+1]
			opcodes[writeTo] = in[inPos]
			pc += 2
			inPos++
		case 4:
			readFrom := opcodes[pc+1]
			return opcodes[readFrom]
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
					return opcodes[pc+1]
				}

				pc += 2
			case 5: // jump-if-true
				var v1, v2 int
				if code%10 == 1 {
					v1 = opcodes[pc+1]
				} else {
					v1 = opcodes[opcodes[pc+1]]
				}

				code /= 10

				if code%10 == 1 {
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
				if code%10 == 1 {
					v1 = opcodes[pc+1]
				} else {
					v1 = opcodes[opcodes[pc+1]]
				}

				code /= 10
				if code%10 == 1 {
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

				if code%10 == 1 {
					v1 = opcodes[pc+1]
				} else {
					v1 = opcodes[opcodes[pc+1]]
				}

				code /= 10

				if code%10 == 1 {
					v2 = opcodes[pc+2]
				} else {
					v2 = opcodes[opcodes[pc+2]]
				}

				if v1 < v2 {
					opcodes[opcodes[pc+3]] = 1
				} else {
					opcodes[opcodes[pc+3]] = 0
				}

				pc += 4
			case 8: // equals
				var v1, v2 int

				if code%10 == 1 {
					v1 = opcodes[pc+1]
				} else {
					v1 = opcodes[opcodes[pc+1]]
				}

				code /= 10

				if code%10 == 1 {
					v2 = opcodes[pc+2]
				} else {
					v2 = opcodes[opcodes[pc+2]]
				}

				if v1 == v2 {
					opcodes[opcodes[pc+3]] = 1
				} else {
					opcodes[opcodes[pc+3]] = 0
				}

				pc += 4
			}
		}
	}

	return 0
}
