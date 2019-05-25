package main

import (
	"bufio"
	"fmt"
	"os"
)

type Euclid struct {
	x, y, q int
}

const maxCapacity = 1024 * 1024 * 10
const bigNum = 1000000007

// Calculate inverse modulo with extended Euclid.
func extendedEuclid(a, b int) Euclid {
	if b == 0 {
		return Euclid{1, 0, a}
	}

	euclid := extendedEuclid(b, a%b)
	x, y, q := euclid.x, euclid.y, euclid.q
	x, y = y, (x - (a/b)*y)
	return Euclid{x, y, q}
}

// Calculate power series recursively.
func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	half := pow(x, n/2)
	res := half * half % bigNum
	if n%2 == 1 {
		res = res * x % bigNum
	}
	return res
}

func main() {
	var (
		t       int
		input   []string
		reverse [1000001]int
	)

	// Read T (Number of test cases)
	fmt.Scanln(&t)

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for i := 0; i < t; i++ {
		scanner.Scan()
		input = append(input, scanner.Text())
	}

	// Pre-calculation of inverse modulo
	for i := 1; i <= 1000000; i++ {
		reverse[i] = extendedEuclid(i, bigNum).x
		for reverse[i] < 0 {
			reverse[i] += bigNum
		}
	}

	// Deal with each test case
	for i := 0; i < t; i++ {
		var n, k, x1, y1, c, d, e1, e2, f int
		fmt.Sscan(input[i], &n, &k, &x1, &y1, &c, &d, &e1, &e2, &f)

		x := make([]int, n)
		x[0] = x1
		y := make([]int, n)
		y[0] = y1
		a := make([]int, n)
		a[0] = (x1 + y1) % f
		for j := 1; j < n; j++ {
			x[j] = (x[j-1]*c + y[j-1]*d + e1) % f
			y[j] = (x[j-1]*d + y[j-1]*c + e2) % f
			a[j] = (x[j] + y[j]) % f
		}

		paramSum := make([]int, n)
		paramSum[n-1] = a[n-1]
		for j := 2; j <= n; j++ {
			paramSum[n-j] = (paramSum[n-j+1] + a[n-j]*j) % bigNum
		}

		powerSum := 0

		for s := 0; s < n; s++ {
			var power int
			if s == 0 {
				power = paramSum[s] * k
			} else {
				mul := (pow(s+1, k+1) - (s + 1)) * reverse[s] % bigNum
				power = paramSum[s] * mul % bigNum
			}
			powerSum = (powerSum + power) % bigNum
		}

		// Print the answer
		fmt.Printf("Case #%d: %d\n", i+1, powerSum)
	}
}
