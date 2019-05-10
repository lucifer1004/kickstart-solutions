package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	const maxCapacity = 1024 * 1024

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
		// Read N (Number of students) & P (Number of players)
		var n, p int
		fmt.Sscan(input[2*i], &n, &p)

		s := strings.Split(input[2*i+1], " ")
		skills := []int{}

		// Read skill of each student
		for j := 0; j < n; j++ {
			skill, _ := strconv.Atoi(s[j])
			skills = append(skills, skill)
		}

		// Sort skill in descending order
		sort.Ints(skills)

		skillSums := []int{0}

		// Calculate accumulative sum
		for j := 0; j < n; j++ {
			skillSums = append(skillSums, skillSums[len(skillSums)-1]+skills[n-j-1])
		}

		// Set original value for the answer
		minHours := -1
		for j := 0; j < n-p+1; j++ {
			// Calculate current skill
			currentSkillSum := skillSums[j+p] - skillSums[j]

			// Calculate target skill
			targetSkillSum := skills[n-j-1] * p

			// Calculate hours required
			hours := targetSkillSum - currentSkillSum

			// Update answer with new minimal
			if hours < minHours || minHours < 0 {
				minHours = hours
			}
		}

		// Print the answer
		fmt.Printf("Case #%d: %d\n", i+1, minHours)
	}
}
