package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type segment struct {
	from, to point
}

// nolint
var (
	in  io.Reader = os.Stdin
	out io.Writer = os.Stdout
)

func main() {
	var wire1, wire2 string

	fmt.Fscanf(in, "%s", &wire1)
	fmt.Fscanf(in, "%s", &wire2)

	wire1Seg := getSegments(wire1)
	wire2Seg := getSegments(wire2)

	fmt.Fprintln(out, closestIntersectionToOrigin(wire1Seg, wire2Seg))
}

func closestIntersectionToOrigin(wire1, wire2 []segment) int {
	dist := int(^uint(0) >> 1)

	for _, s1 := range wire1 {
		for _, s2 := range wire2 {
			pt, err := intersection(s1, s2)
			if err == nil {
				if d := pt.x + pt.y; d < dist {
					dist = d
				}
			}
		}
	}

	return dist
}

func intersection(s1, s2 segment) (point, error) {
	var horiz, verti segment

	if s1.from.y == s1.to.y { // s1 is horizontal
		horiz = s1

		if s2.from.x != s2.to.x {
			return point{}, fmt.Errorf("not perpendicular")
		}

		verti = s2
	} else if s2.from.y == s2.to.y { // s2 is horizontal
		horiz = s2

		if s1.from.x != s1.to.x {
			return point{}, fmt.Errorf("not perpendicular")
		}
		verti = s1
	} else {
		return point{}, fmt.Errorf("neither horizontal")
	}

	leftX, rightX := horiz.from.x, horiz.to.x
	if leftX > rightX {
		leftX, rightX = rightX, leftX
	}

	if verti.from.x >= leftX && verti.from.x <= rightX {
		bottomY, topY := verti.from.y, verti.to.y
		if bottomY > topY {
			bottomY, topY = topY, bottomY
		}

		if horiz.from.y >= bottomY && horiz.from.y <= topY {
			return point{verti.from.x, horiz.from.y}, nil
		}
	}

	return point{}, fmt.Errorf("don't intersect")
}

func getSegments(wire string) []segment {
	var ans []segment

	lastPoint := point{0, 0}

	for _, move := range strings.Split(wire, ",") {
		move = strings.TrimSpace(move)
		magnitude, _ := strconv.Atoi(move[1:])

		var to point

		switch rune(move[0]) {
		case 'R':
			to = point{lastPoint.x + magnitude, lastPoint.y}
			ans = append(ans, segment{lastPoint, to})
		case 'D':
			to = point{lastPoint.x, lastPoint.y - magnitude}
			ans = append(ans, segment{lastPoint, to})
		case 'L':
			to = point{lastPoint.x - magnitude, lastPoint.y}
			ans = append(ans, segment{lastPoint, to})
		case 'U':
			to = point{lastPoint.x, lastPoint.y + magnitude}
			ans = append(ans, segment{lastPoint, to})
		}

		lastPoint = to
	}

	return ans
}
