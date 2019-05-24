package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	sum, prefix int
}

type SegmentTree []Node

func sumNode(l, r Node) Node {
	sum := l.sum + r.sum
	prefix := l.prefix
	rightMax := l.sum + r.prefix
	if prefix < rightMax {
		prefix = rightMax
	}
	return Node{
		sum:    sum,
		prefix: prefix,
	}
}

func buildTree(nums []int) SegmentTree {
	length := len(nums) - 1
	count := uint(0)
	for length > 0 {
		length /= 2
		count++
	}
	length = 1 << count
	tree := make([]Node, 2*length)
	buildTreeUtil(tree, nums, 0, len(nums)-1, 0)
	return tree
}

func buildTreeUtil(tree SegmentTree, nums []int, l, r, index int) Node {
	if l == r {
		tree[index] = Node{
			sum:    nums[l],
			prefix: nums[l],
		}
	} else {
		mid := (l + r) / 2
		left := buildTreeUtil(tree, nums, l, mid, index*2+1)
		right := buildTreeUtil(tree, nums, mid+1, r, index*2+2)
		tree[index] = sumNode(left, right)
	}

	return tree[index]
}

func updateTree(tree SegmentTree, length, pos, val int) {
	updateTreeUtil(tree, 0, length-1, 0, pos, val)
}

func updateTreeUtil(tree SegmentTree, l, r, index, pos, val int) Node {
	if l == r {
		tree[index] = Node{
			sum:    val,
			prefix: val,
		}
	} else {
		mid := (l + r) / 2
		left := tree[index*2+1]
		right := tree[index*2+2]
		if pos <= mid {
			left = updateTreeUtil(tree, l, mid, index*2+1, pos, val)
		} else {
			right = updateTreeUtil(tree, mid+1, r, index*2+2, pos, val)
		}
		tree[index] = sumNode(left, right)
	}

	return tree[index]
}

func queryTree(tree SegmentTree, length, begin int) int {
	return queryTreeUtil(tree, begin, length-1, 0, length-1, 0).prefix
}

func queryTreeUtil(tree SegmentTree, begin, end, l, r, index int) Node {
	res := Node{}
	if l >= begin && r <= end {
		res = tree[index]
	} else {
		mid := (l + r) / 2
		if begin <= mid {
			res = sumNode(res, queryTreeUtil(tree, begin, end, l, mid, index*2+1))
		}
		if end >= mid+1 {
			res = sumNode(res, queryTreeUtil(tree, begin, end, mid+1, r, index*2+2))
		}
	}
	return res
}

func main() {
	const maxCapacity = 1024 * 1024

	var (
		t     int
		n, s  []int
		input []string
	)

	// Read T (Number of test cases)
	fmt.Scanln(&t)

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for i := 0; i < t; i++ {
		var nI, sI int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &nI, &sI)
		n = append(n, nI)
		s = append(s, sI)

		scanner.Scan()
		input = append(input, scanner.Text())
	}

	// Deal with each test case
	for i := 0; i < t; i++ {
		a := make([]int, n[i])
		v := make([]int, n[i])
		str := strings.Split(input[i], " ")
		dict := map[int][]int{}
		dictIndex := map[int]int{}

		for j := 0; j < n[i]; j++ {
			a[j], _ = strconv.Atoi(str[j])
			dict[a[j]] = append(dict[a[j]], j)
			if len(dict[a[j]]) < s[i]+1 {
				v[j] = 1
			}
			if len(dict[a[j]]) == s[i]+1 {
				v[j] = -s[i]
			}
		}

		tree := buildTree(v)

		maxBooks := 0

		for j := 0; j < n[i]; j++ {

			if j >= 1 {
				dictIndex[a[j-1]]++
				length := len(dict[a[j-1]])
				count := dictIndex[a[j-1]]
				if length-count >= s[i] {
					k1 := dict[a[j-1]][count+s[i]-1]
					updateTree(tree, n[i], k1, 1)
				}
				if length-count >= s[i]+1 {
					k2 := dict[a[j-1]][count+s[i]]
					updateTree(tree, n[i], k2, -s[i])
				}
			}

			books := queryTree(tree, n[i], j)

			if books > maxBooks {
				maxBooks = books
			}
		}

		// Print the answer
		fmt.Printf("Case #%d: %d\n", i+1, maxBooks)
	}
}
