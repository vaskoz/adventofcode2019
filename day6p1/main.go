package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// nolint
var (
	in  io.Reader = os.Stdin
	out io.Writer = os.Stdout
)

func main() {
	adjList := make(map[string]map[string]struct{})
	orbit := ""

	for _, err := fmt.Fscanf(in, "%s", &orbit); err != io.EOF; _, err = fmt.Fscanf(in, "%s", &orbit) {
		parts := strings.Split(orbit, ")")
		if _, exists := adjList[parts[0]]; !exists {
			adjList[parts[0]] = make(map[string]struct{})
		}

		adjList[parts[0]][parts[1]] = struct{}{}
	}

	ans := 0
	discovered := make(map[string]struct{})
	discovered["COM"] = struct{}{}
	q := make([]string, 0, 1)
	q = append(q, "COM")

	dist := 1

	for len(q) != 0 {
		nextQ := make([]string, 0)

		for len(q) != 0 {
			v := q[0]
			q = q[1:]

			for v2 := range adjList[v] {
				if _, done := discovered[v2]; !done {
					ans += dist
					discovered[v2] = struct{}{}

					nextQ = append(nextQ, v2)
				}
			}
		}

		q = nextQ
		dist++
	}

	fmt.Fprintln(out, ans)
}
