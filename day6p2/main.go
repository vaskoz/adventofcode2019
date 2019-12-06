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
	adjList := make(map[string]string)
	orbit := ""

	for _, err := fmt.Fscanf(in, "%s", &orbit); err != io.EOF; _, err = fmt.Fscanf(in, "%s", &orbit) {
		parts := strings.Split(orbit, ")")
		adjList[parts[1]] = parts[0]
	}

	var youToCom, sanToCom []string

	for v := "YOU"; v != ""; v = adjList[v] {
		youToCom = append(youToCom, v)
	}

	for v := "SAN"; v != ""; v = adjList[v] {
		sanToCom = append(sanToCom, v)
	}

	lastCommon := 0

	for ; sanToCom[len(sanToCom)-1-lastCommon] == youToCom[len(youToCom)-1-lastCommon]; lastCommon++ {
	}

	youToCom = youToCom[:len(youToCom)-lastCommon]
	sanToCom = sanToCom[:len(sanToCom)-lastCommon]

	fmt.Fprintln(out, len(youToCom)+len(sanToCom)-2)
}
