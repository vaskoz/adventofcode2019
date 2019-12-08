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

	for a := 5; a < 10; a++ {
		used[a] = struct{}{}
		for b := 5; b < 10; b++ {
			if _, use := used[b]; use {
				continue
			}
			used[b] = struct{}{}
			for c := 5; c < 10; c++ {
				if _, use := used[c]; use {
					continue
				}
				used[c] = struct{}{}
				for d := 5; d < 10; d++ {
					if _, use := used[d]; use {
						continue
					}
					used[d] = struct{}{}
					for e := 5; e < 10; e++ {
						if _, use := used[e]; use {
							continue
						}
						inA := make(chan int, 2)
						inA <- a
						inA <- 0
						inB := make(chan int, 2)
						inB <- b
						inC := make(chan int, 2)
						inC <- c
						inD := make(chan int, 2)
						inD <- d
						inE := make(chan int, 2)
						inE <- e

						outA := run(inA)
						outB := run(inB)
						outC := run(inC)
						outD := run(inD)
						outE := run(inE)
						thrust := 0

						for {
							if data, cl := <-outA; !cl {
								break
							} else {
								inB <- data
								thrust = data
							}
							if data, cl := <-outB; !cl {
								break
							} else {
								inC <- data
								thrust = data
							}
							if data, cl := <-outC; !cl {
								break
							} else {
								inD <- data
								thrust = data
							}
							if data, cl := <-outD; !cl {
								break
							} else {
								inE <- data
								thrust = data
							}
							if data, cl := <-outE; !cl {
								break
							} else {
								inA <- data
								thrust = data
							}
						}
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

func run(in <-chan int) chan int {
	ch := make(chan int)
	f, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(f)
	f.Close()
	prog := string(data)
	prog = strings.TrimSpace(prog)

	strs := strings.Split(prog, ",")
	opcodes := make([]int, len(strs))

	for i, s := range strs {
		res, _ := strconv.Atoi(s)
		opcodes[i] = res
	}

	go func() {

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
				opcodes[writeTo] = <-in
				pc += 2
			case 4:
				readFrom := opcodes[pc+1]
				ch <- opcodes[readFrom]
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
						ch <- opcodes[pc+1]
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

		close(ch)
	}()

	return ch
}
