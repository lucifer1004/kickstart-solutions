package main

import (
	"fmt"
)

func main() {
	var t int

	// Read T (Number of test cases)
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		var (
			a, b int
			n    int
			res  string
		)
		fmt.Scanln(&a, &b)
		a++
		fmt.Scanln(&n)

		for a <= b {
			mid := a + (b-a)/2
			fmt.Println(mid)
			fmt.Scanln(&res)
			if res == "TOO_BIG" {
				b = mid - 1
			} else if res == "TOO_SMALL" {
				a = mid + 1
			} else if res == "CORRECT" {
				break
			}
		}
	}
}
