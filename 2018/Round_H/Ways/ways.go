package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const maxCapacity = 1024 * 1024 * 10
const maxN = 200001
const modulo = 1000000007

var scanner *bufio.Scanner
var writer *bufio.Writer
var fac, two, bMod [maxN]int

func quickMod(a, b int) int {
	ans := 1
	for b > 0 {
		if b&1 == 1 {
			ans = ans * a % modulo
			b--
		}
		b >>= 1
		a = a * a % modulo
	}
	return ans
}

func g(n, m int) int {
	return fac[2*n-m] * two[m] % modulo
}

func sig(n int) int {
	if n%2 == 0 {
		return 1
	}
	return -1
}

func ioPrepare() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
}

func solve(caseNum int) {
	var (
		n, m int
		c    []int
	)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n, &m)
	ans := fac[2*n]

	// calculate combinations
	c = make([]int, m+1)
	c[0] = 1

	for i := 1; i <= m; i++ {
		c[i] = c[i-1] * (m - i + 1) % modulo
		c[i] = c[i] * bMod[i] % modulo
	}

	for i := 1; i <= m; i++ {
		ans = (ans + sig(i)*c[i]*g(n, i)) % modulo
	}

	outputStr := fmt.Sprintf("Case #%d: %d\n", caseNum+1, (ans+modulo)%modulo)

	writer.WriteString(outputStr)
}

func main() {
	fac[0] = 1
	two[0] = 1
	bMod[0] = 0

	// Pre-calculate factorials, powers of 2 and modular inverses
	for i := 1; i < maxN; i++ {
		fac[i] = i * fac[i-1] % modulo
		two[i] = 2 * two[i-1] % modulo
		bMod[i] = quickMod(i, modulo-2)
	}

	ioPrepare()
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		solve(i)
	}

	writer.Flush()
}
