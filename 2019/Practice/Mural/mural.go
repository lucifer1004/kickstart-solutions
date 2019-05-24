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

	var (
		t     int
		input []string
	)

	// Read T (Number of test cases)
	fmt.Scanln(&t)

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for i := 0; i < 2*t; i++ {
		scanner.Scan()
		input = append(input, scanner.Text())
	}

	// Deal with each test case
	for i := 0; i < t; i++ {
		// Read N (Number of sections)
		var n int
		fmt.Sscan(input[2*i], &n)
		s := strings.Split(input[2*i+1], "")
		beauty := make([]int, n)

		for j := 0; j < n; j++ {
			beauty[j], _ = strconv.Atoi(s[j])
		}

		accumulatedBeauty := make([]int, n+1)
		accumulatedBeauty[0] = 0

		for j := 0; j < n; j++ {
			accumulatedBeauty[j+1] = accumulatedBeauty[j] + beauty[j]
		}

		length := (n + 1) / 2
		maxBeauty := 0

		for j := 0; j < n-length+1; j++ {
			currentBeauty := accumulatedBeauty[j+length] - accumulatedBeauty[j]
			if currentBeauty > maxBeauty {
				maxBeauty = currentBeauty
			}
		}

		// Print the answer
		fmt.Printf("Case #%d: %d\n", i+1, maxBeauty)
	}
}
