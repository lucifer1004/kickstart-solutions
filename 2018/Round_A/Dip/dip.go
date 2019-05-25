package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int

	// Read T (Number of test cases)
	fmt.Scanln(&t)

	// Deal with each test case
	for i := 0; i < t; i++ {
		var (
			n, k                    int
			v                       float64
			values, sum, mean, best []float64
		)
		fmt.Scanln(&n, &k)
		for j := 0; j < n; j++ {
			fmt.Scan(&v)
			values = append(values, v)
		}
		sort.Float64s(values)

		for j := n - 1; j >= 0; j-- {
			if j == n-1 {
				sum = append(sum, values[j])
				mean = append(mean, values[j])
			} else {
				sum = append(sum, values[j]+sum[len(sum)-1])
				mean = append(mean, sum[len(sum)-1]/float64(len(sum)))
			}
		}

		best = append(best, mean[n-1])
		searchLimit := n - 2

		for j := 1; j <= k; j++ {
			currentBest := mean[n-1]
			bestWatershed := n - 1
			for watershed := 0; watershed <= searchLimit; watershed++ {
				p := float64(watershed+1) / float64(n)
				newValue := p*mean[watershed] + (float64(1)-p)*best[j-1]
				if newValue > currentBest {
					currentBest = newValue
					bestWatershed = watershed
				}
			}
			best = append(best, currentBest)
			searchLimit = bestWatershed
			// fmt.Println("Redip:", j, bestWatershed, currentBest)
		}

		// Print the answer
		fmt.Printf("Case #%d: %f\n", i+1, best[k])
	}
}
