package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State struct {
	dist    int
	visited map[int]bool
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
	for i := 0; i < 3*t; i++ {
		scanner.Scan()
		input = append(input, scanner.Text())
	}

	// Deal with each test case
	for i := 0; i < t; i++ {
		var (
			n, k int
			d, c []int
		)

		fmt.Sscan(input[3*i], &n, &k)
		distanceStr := strings.Split(input[3*i+1], " ")
		colorStr := strings.Split(input[3*i+2], " ")

		for j := 0; j < n; j++ {
			distance, _ := strconv.Atoi(distanceStr[j])
			color, _ := strconv.Atoi(colorStr[j])
			d = append(d, distance)
			c = append(c, color)
		}

		row := []int{0}
		row = append(row, d...)
		mat := [][]int{row}
		for j := 0; j < n; j++ {
			row := []int{d[j]}
			for s := 0; s < n; s++ {
				if s == j {
					row = append(row, 0)
					continue
				}
				if c[j] == c[s] {
					dist := d[s] - d[j]
					if dist < 0 {
						dist = -dist
					}
					row = append(row, dist)
				} else {
					row = append(row, d[s]+d[j])
				}
			}
			mat = append(mat, row)
		}

		val := [][]State{}
		for j := 1; j <= k; j++ {
			value := []State{}
			for s := 1; s <= n; s++ {
				visited := map[int]bool{}
				if j == 1 {
					visited[s] = true
					state := State{
						dist:    mat[0][s],
						visited: visited,
					}
					value = append(value, state)
				} else {
					minDist := -1
					minDistIndex := -1
					for w := 1; w <= n; w++ {
						if w == s || val[j-2][w-1].visited[s] || val[j-2][w-1].dist == -1 {
							continue
						} else {
							dist := val[j-2][w-1].dist + mat[w][s]
							if dist < minDist || minDist < 0 {
								minDist = dist
								minDistIndex = w - 1
							}
						}
					}

					if minDistIndex >= 0 {
						for k, v := range val[j-2][minDistIndex].visited {
							visited[k] = v
						}
						visited[s] = true
					}

					state := State{
						dist:    minDist,
						visited: visited,
					}

					value = append(value, state)
				}
			}
			val = append(val, value)
		}

		minDist := -1
		for j := 1; j <= n; j++ {
			if val[k-1][j-1].dist < minDist || minDist < 0 {
				minDist = val[k-1][j-1].dist
			}
		}

		outputStr := fmt.Sprintf("Case #%d: %d\n", i+1, minDist)
		writer.WriteString(outputStr)
	}
	writer.Flush()
}
