package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Stone ...
// The magic stone
type Stone struct {
	timeToEat     int
	originalValue int
	lossPerSecond int
}

// Stones ...
// The magic stones
type Stones []Stone

// Len ...
func (s Stones) Len() int {
	return len(s)
}

// Swap ...
func (s Stones) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less ...
func (s Stones) Less(i, j int) bool {
	return s[i].timeToEat*s[j].lossPerSecond < s[j].timeToEat*s[i].lossPerSecond
}

func main() {
	const maxCapacity = 1024 * 1024 * 10

	var (
		t     int
		n     []int
		input []string
	)

	// Read T (Number of test cases)
	fmt.Scanln(&t)

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for i := 0; i < t; i++ {
		var nI int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &nI)
		n = append(n, nI)
		for j := 0; j < nI; j++ {
			scanner.Scan()
			input = append(input, scanner.Text())
		}
	}

	line := 0

	// Deal with each test case
	for i := 0; i < t; i++ {
		var (
			s, e, l int
			stones  Stones
			v       [][]int
		)

		for j := 0; j < n[i]; j++ {
			fmt.Sscan(input[line], &s, &e, &l)
			stone := Stone{
				timeToEat:     s,
				originalValue: e,
				lossPerSecond: l,
			}
			stones = append(stones, stone)
			line++
		}

		sort.Sort(stones)

		totalTime := 0

		for _, value := range stones {
			totalTime += value.timeToEat
		}

		for j := 0; j <= totalTime; j++ {
			z := make([]int, n[i]+1)
			v = append(v, z)
		}

		maxValue := 0

		for j := 1; j <= totalTime; j++ {
			for k := 1; k <= n[i]; k++ {
				s := stones[k-1].timeToEat
				e := stones[k-1].originalValue
				l := stones[k-1].lossPerSecond

				v[j][k] = v[j][k-1]

				if j < s {
					continue
				}

				currentValue := e - (j-s)*l
				if currentValue < 0 {
					currentValue = 0
				}

				if v[j][k] < v[j-s][k-1]+currentValue {
					v[j][k] = v[j-s][k-1] + currentValue
				}

				if v[j][k] > maxValue {
					maxValue = v[j][k]
				}
			}
		}

		// Print the answer
		fmt.Printf("Case #%d: %d\n", i+1, maxValue)
	}
}
