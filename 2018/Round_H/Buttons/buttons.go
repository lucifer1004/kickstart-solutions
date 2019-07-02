package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type strArr []string

func (s strArr) Len() int {
	return len(s)
}
func (s strArr) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
func (s strArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

const maxCapacity = 1024 * 1024 * 10

var scanner *bufio.Scanner
var writer *bufio.Writer

func twoPower(n int) int64 {
	if n < 0 {
		panic("negative power is not supported yet")
	}
	if n == 0 {
		return int64(1)
	}
	half := n / 2
	ans := twoPower(half) * twoPower(half)
	if n%2 == 0 {
		return ans
	}
	return ans * 2
}

func ioPrepare() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
}

func solve(caseNum int) {
	var (
		n, p   int
		ans    int64
		forbid strArr
	)

	forbid = strArr{}

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n, &p)

	ans = twoPower(n)

	for i := 0; i < p; i++ {
		scanner.Scan()
		forbid = append(forbid, scanner.Text())
	}
	sort.Sort(forbid)

	dup := map[int]bool{}

	for i := 0; i < p-1; i++ {
		for j := i + 1; j < p; j++ {
			if strings.Index(forbid[j], forbid[i]) == 0 {
				dup[j] = true
			}
		}
	}

	for i := 0; i < p; i++ {
		if !dup[i] {
			ans -= twoPower(n - len(forbid[i]))
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
