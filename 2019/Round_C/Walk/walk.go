package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	var input []string
	const maxCapacity = 1024 * 1024 * 10

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < 2*t; i++ {
		scanner.Scan()
		input = append(input, scanner.Text())
	}

	// Deal with each test case
	for i := 0; i < t; i++ {
		var n, r, c, sr, sc int
		var (
			instructions []string
		)

		fmt.Sscan(input[2*i], &n, &r, &c, &sr, &sc)
		instructions = strings.Split(input[2*i+1], "")

		visited := map[Point]bool{}
		directions := map[string]Point{
			"W": Point{x: 0, y: -1},
			"E": Point{x: 0, y: 1},
			"N": Point{x: -1, y: 0},
			"S": Point{x: 1, y: 0},
		}

		visited[Point{x: sr, y: sc}] = true

		for j := 0; j < len(instructions); j++ {
			var left, right, mid int
			ins := instructions[j]
			switch ins {
			case "W":
				left = 1
				right = sc - 1
			case "E":
				left = 1
				right = c - sc
			case "N":
				left = 1
				right = sr - 1
			case "S":
				left = 1
				right = r - sr
			}
			for left < right {
				mid = (left + right) / 2
				if visited[Point{x: sr + mid*directions[ins].x, y: sc + mid*directions[ins].y}] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			sr += left * directions[ins].x
			sc += left * directions[ins].y
			visited[Point{x: sr, y: sc}] = true
		}

		outputStr := fmt.Sprintf("Case #%d: %d %d\n", i+1, sr, sc)
		writer.WriteString(outputStr)
		writer.Flush()
	}
}
