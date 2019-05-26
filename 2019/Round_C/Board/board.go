package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const maxCapacity = 1024 * 1024 * 10

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	line := 0

	// Deal with each test case
	for i := 0; i < t; i++ {
		var r, c, diff int
		scanner.Scan()
		fmt.Sscanln(scanner.Text(), &r, &c, &diff)

		row := []int{}
		board := [][]int{}
		for j := 0; j < r; j++ {
			scanner.Scan()
			thickStr := strings.Split(scanner.Text(), " ")
			for k := 0; k < c; k++ {
				thick, _ := strconv.Atoi(thickStr[k])
				row = append(row, thick)
			}
			board = append(board, row)
			row = []int{}
			line++
		}

		boardRowValue := [][]int{}
		for j := 0; j < r; j++ {
			rowValue := []int{}
			for k := 0; k < c; k++ {
				min := board[j][k]
				max := board[j][k]
				length := 1
				for s := k + 1; s < c; s++ {
					value := board[j][s]
					if value > max {
						max = value
					}
					if value < min {
						min = value
					}
					if max-min <= diff {
						length++
					} else {
						break
					}
				}
				rowValue = append(rowValue, length)
			}
			boardRowValue = append(boardRowValue, rowValue)
		}

		maxArea := 0
		for j := 0; j < c; j++ {
			for k := 0; k < r; k++ {
				min := -1
				for s := k; s < r; s++ {
					newValue := boardRowValue[s][j]
					if newValue < min || min < 0 {
						min = newValue
					}
					area := min * (s - k + 1)
					if area > maxArea {
						maxArea = area
					}
				}
			}
		}

		fmt.Printf("Case #%d: %d\n", i+1, maxArea)
	}
}
