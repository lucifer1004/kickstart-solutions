package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const maxCapacity = 1024 * 1024

	var (
		t     int
		n, s  []int
		input []string
	)

	// Read T (Number of test cases)
	fmt.Scanln(&t)

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for i := 0; i < t; i++ {
		var nI, sI int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &nI, &sI)
		n = append(n, nI)
		s = append(s, sI)

		scanner.Scan()
		input = append(input, scanner.Text())
	}

	// Deal with each test case
	for i := 0; i < t; i++ {
		a := make([]int, n[i])
		str := strings.Split(input[i], " ")

		for j := 0; j < n[i]; j++ {
			a[j], _ = strconv.Atoi(str[j])
		}

		maxBooks := 0

		for j := 0; j < n[i]; j++ {
			dict := map[int]int{}
			books := 0
			for k := j; k < n[i]; k++ {
				if dict[a[k]] < s[i] {
					books++
					if books > maxBooks {
						maxBooks = books
					}
				}
				if dict[a[k]] == s[i] {
					books -= s[i]
				}
				dict[a[k]]++
			}
		}

		// Print the answer
		fmt.Printf("Case #%d: %d\n", i+1, maxBooks)
	}
}
