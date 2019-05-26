package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const modulo = 1000000007
const maxCapacity = 1024 * 1024 * 10

var f, wf *os.File
var scanner *bufio.Scanner
var writer *bufio.Writer

func ioPrepare() {
	var err error

	f, err = os.Open("./large.in")
	if err != nil {
		panic("Cannot open input file")
	}
	scanner = bufio.NewScanner(f)
	wf, err = os.Create("./large.out")
	if err != nil {
		panic("Cannot create output file")
	}
	writer = bufio.NewWriter(wf)

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
}

func fastMul(x, y int) int {
	ans := 0
	for y > 0 {
		if y&1 == 1 {
			y--
			ans = (ans + x) % modulo
		}
		y >>= 1
		x = (x + x) % modulo
	}
	return ans
}

func solve(caseNum int) {
	var (
		r, c int
	)
	scanner.Scan()
	fmt.Sscanln(scanner.Text(), &r, &c)

	if r > c {
		r, c = c, r
	}

	ans := 0
	rMulC := r * c % modulo
	rAddC := (r + c) % modulo

	for i := 1; i < r; i++ {
		res1 := rAddC * i % modulo
		res2 := i * i % modulo
		res3 := rMulC * i % modulo
		res4 := res1 * i % modulo
		res5 := res2 * i % modulo
		res6 := (res3 - res4) % modulo
		res7 := (res6 + res5) % modulo
		ans = (ans + res7) % modulo
	}

	if ans < 0 {
		ans = ans + modulo
	}

	// Print the answer
	outputStr := fmt.Sprintf("Case #%d: %d\n", caseNum+1, ans)
	writer.WriteString(outputStr)
	writer.Flush()
}

func main() {
	ioPrepare()

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		solve(i)
	}
}
