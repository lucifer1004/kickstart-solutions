package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Word [26]int

func main() {
	const maxCapacity = 1024 * 1024 * 10

	// // For local test
	// f, _ := os.Open("./visible_test.in")
	// scanner := bufio.NewScanner(f)
	// wf, _ := os.Create("./visible_test.out")
	// writer := bufio.NewWriter(wf)

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	// Deal with each test case
	for i := 0; i < t; i++ {
		var (
			l, n, a, b, c, d, x1, x2 int
			s1, s2, s                string
			letterCount              [][26]int
			buffer                   bytes.Buffer
		)
		scanner.Scan()
		l, _ = strconv.Atoi(scanner.Text())

		scanner.Scan()
		words := strings.Split(scanner.Text(), " ")
		hashDict := map[string][]Word{}
		lengthDict := map[int]bool{}
		remaining := l

		for _, value := range words {
			length := len(value)
			lengthDict[length] = true

			count := [26]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			for j := 1; j < length-1; j++ {
				count[int(value[j])-97]++
			}
			word := count
			hashIndex := string(value[0]) + strconv.Itoa(length) + string(value[length-1])
			hashDict[hashIndex] = append(hashDict[hashIndex], word)
		}

		scanner.Scan()
		params := strings.Split(scanner.Text(), " ")
		s1 = params[0]
		s2 = params[1]
		n, _ = strconv.Atoi(params[2])
		a, _ = strconv.Atoi(params[3])
		b, _ = strconv.Atoi(params[4])
		c, _ = strconv.Atoi(params[5])
		d, _ = strconv.Atoi(params[6])
		x1 = int(s1[0])
		x2 = int(s2[0])
		buffer.WriteString(s1)
		buffer.WriteString(s2)
		for j := 3; j <= n; j++ {
			x1, x2 = x2, (a*x2+b*x1+c)%d
			buffer.WriteString(string(97 + x2%26))
		}
		s = buffer.String()

		letterCount = append(letterCount, [26]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		for j := 0; j < n; j++ {
			nextCount := [26]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			for k := 0; k < 26; k++ {
				nextCount[k] = letterCount[j][k]
			}
			nextCount[int(s[j])-97]++
			letterCount = append(letterCount, nextCount)
		}

		for j := 0; j < len(s)-1; j++ {
			for k := range lengthDict {
				if j+k > len(s) {
					continue
				}
				hashIndex := string(s[j]) + strconv.Itoa(k) + string(s[j+k-1])
				hashResult := hashDict[hashIndex]
				if len(hashResult) == 0 {
					continue
				}

				count := [26]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
				for w := 0; w < 26; w++ {
					count[w] = letterCount[j+k-1][w] - letterCount[j+1][w]
				}

				used := map[int]bool{}

				for w := 0; w < len(hashResult); w++ {
					word := hashResult[w]
					if word == count {
						// fmt.Println(j, k, s[j:j+k], hashIndex)
						remaining--

						if len(hashResult) == 1 {
							hashDict[hashIndex] = nil
						} else {
							used[w] = true
						}
					}
				}

				if len(used) > 0 {
					newHash := []Word{}
					for w := 0; w < len(hashResult); w++ {
						if !used[w] {
							newHash = append(newHash, hashResult[w])
						}
					}

					hashDict[hashIndex] = newHash
				}

				if remaining == 0 {
					break
				}
			}
			if remaining == 0 {
				break
			}
		}

		// Print the answer
		outputStr := fmt.Sprintf("Case #%d: %d\n", i+1, l-remaining)
		writer.WriteString(outputStr)
		writer.Flush()
	}
}
