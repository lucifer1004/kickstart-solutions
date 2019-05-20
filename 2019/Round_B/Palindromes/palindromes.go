package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const maxCapacity = 1024 * 1024 * 10

	var (
		t     int
		n     []int
		q     []int
		input []string
	)

	// Read T (Number of test cases)
	fmt.Scanln(&t)

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for i := 0; i < t; i++ {
		var (
			nI, qI int
		)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &nI, &qI)
		n = append(n, nI)
		q = append(q, qI)
		for j := 0; j < qI+1; j++ {
			scanner.Scan()
			input = append(input, scanner.Text())
		}
	}

	line := 0

	// Deal with each test case
	for i := 0; i < t; i++ {
		var letterCount [][26]int
		var tmpCount [26]int
		s := input[line]
		possibleCount := 0
		letterCount = append(letterCount, tmpCount)
		for j := 0; j < len(s); j++ {
			tmpCount := letterCount[j]
			letter := int(s[j]) - int('A')
			tmpCount[letter]++
			letterCount = append(letterCount, tmpCount)
		}
		line++
		for j := 0; j < q[i]; j++ {
			var l, r, oddCount int
			fmt.Sscan(input[line], &l, &r)
			for k := 0; k < 26; k++ {
				if (letterCount[r][k]-letterCount[l-1][k])%2 == 1 {
					oddCount++
				}
			}
			if oddCount <= 1 {
				possibleCount++
			}
			line++
		}

		// Print the answer
		fmt.Printf("Case #%d: %d\n", i+1, possibleCount)
	}
}
