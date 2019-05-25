package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func partialConvert(a []int, l, r int) int64 {
	res := int64(0)
	for l <= r {
		res = res*10 + int64(a[l])
		l++
	}
	return res
}

func main() {
	var t int
	tenPower := make([]int64, 17)
	tenPower[0] = int64(1)

	// Read T (Number of test cases)
	fmt.Scanln(&t)

	for i := 1; i < 16; i++ {
		tenPower[i] = tenPower[i-1] * 10
	}

	// Deal with each test case
	for i := 0; i < t; i++ {
		var n int64
		fmt.Scanln(&n)
		nStr := strconv.FormatInt(n, 10)
		chars := strings.Split(nStr, "")
		digits := make([]int, len(chars))
		for j := range digits {
			digits[j], _ = strconv.Atoi(chars[j])
		}

		minAttempts := int64(0)

		for j := range digits {
			if digits[j]%2 == 0 {
				continue
			} else {
				current := partialConvert(digits, j, len(digits)-1)
				if digits[j] == 9 {
					lower := []int{digits[j] - 1}
					for k := 0; k < len(digits)-j-1; k++ {
						lower = append(lower, 8)
					}
					target := partialConvert(lower, 0, len(lower)-1)
					minAttempts = current - target
				} else {
					upper := []int{digits[j] + 1}
					lower := []int{digits[j] - 1}
					for k := 0; k < len(digits)-j-1; k++ {
						lower = append(lower, 8)
						upper = append(upper, 0)
					}
					lowerTarget := partialConvert(lower, 0, len(lower)-1)
					upperTarget := partialConvert(upper, 0, len(upper)-1)
					minAttempts = int64(math.Min(float64(current-lowerTarget), float64(upperTarget-current)))
				}
				break
			}
		}

		// Print the answer
		fmt.Printf("Case #%d: %d\n", i+1, minAttempts)
	}
}
