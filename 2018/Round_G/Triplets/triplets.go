package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func binarySearchFirstAndLast(arr []int, val int) [2]int {
	first := -1
	last := -1

	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == val {
			if mid == 0 || arr[mid-1] < val {
				first = mid
			}
			r = mid - 1
		} else if arr[mid] < val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	l = 0
	r = len(arr) - 1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == val {
			if mid == len(arr)-1 || arr[mid+1] > val {
				last = mid
			}
			l = mid + 1
		} else if arr[mid] < val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return [2]int{first, last}
}

func solve(caseNum int) {
	var (
		n int
	)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	d := []int{}
	z := 0
	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), " ") {
		t, _ := strconv.Atoi(v)
		if t == 0 {
			z++
		} else {
			d = append(d, t)
		}
	}
	max := 0
	if len(d) > 0 {
		sort.Ints(d)
		max = d[len(d)-1]
	}
	ans := z*(z-1)/2*len(d) + z*(z-1)*(z-2)/6

	for i := 0; i < len(d); i++ {
		if d[i]*d[i] > max {
			break
		}
		for j := i + 1; j < len(d); j++ {
			target := d[i] * d[j]
			if target > max {
				break
			}
			find := binarySearchFirstAndLast(d[j+1:], target)
			if find[0] != -1 {
				ans += find[1] - find[0] + 1
			}
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
