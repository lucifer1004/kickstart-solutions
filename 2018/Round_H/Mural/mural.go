package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxCapacity = 1024 * 1024 * 10

var scanner *bufio.Scanner
var writer *bufio.Writer

func ioPrepare() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
}

func solve(caseNum int) {
	var (
		n int
		b []int
	)
	b = []int{}

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	bs := strings.Split(scanner.Text(), "")

	for _, v := range bs {
		digit, _ := strconv.Atoi(string(v))
		b = append(b, digit)
	}
	maxLen := (len(b) + 1) / 2
	sum := 0
	for i := 0; i < maxLen; i++ {
		sum += b[i]
	}
	ans := sum
	for i := 1; i <= len(b)-maxLen; i++ {
		sum -= b[i-1]
		sum += b[i-1+maxLen]
		if sum > ans {
			ans = sum
		}
	}

	outputStr := fmt.Sprintf("Case #%d: %d\n", caseNum+1, ans)
	writer.WriteString(outputStr)
}

func main() {
	ioPrepare()
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		solve(i)
	}

	writer.Flush()
}
