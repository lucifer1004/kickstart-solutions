package rangetree1d

type rangeTree struct {
	val          int
	ltree, rtree *rangeTree
}

func build(vals []int) rangeTree {
	if len(vals) == 0 {
		panic("Cannot initialize an empty range tree")
	}
	tree := rangeTree{val: vals[0]}
	for i := 1; i < len(vals); i++ {
		insert(&tree, vals[i])
	}
	return tree
}

func insert(tree *rangeTree, val int) {
	if val == tree.val {
		return
	}
	if val < tree.val {
		if tree.ltree != nil {
			insert(tree.ltree, val)
		} else {
			tree.ltree = &rangeTree{
				val: val,
			}
		}
	} else {
		if tree.rtree != nil {
			insert(tree.rtree, val)
		} else {
			tree.rtree = &rangeTree{
				val: val,
			}
		}
	}
}

func delete(tree *rangeTree, val int) *rangeTree {
	if val == tree.val {
		if tree.ltree == nil {
			return tree.rtree
		}
		if tree.rtree == nil {
			return tree.ltree
		}
		node := tree.rtree
		for node.ltree != nil {
			node = node.ltree
		}
		node.ltree = tree.ltree
		return node
	}
	if val < tree.val {
		if tree.ltree != nil {
			tree.ltree = delete(tree.ltree, val)
		}
	} else {
		if tree.rtree != nil {
			tree.rtree = delete(tree.rtree, val)
		}
	}
	return tree
}

func traverse(tree rangeTree) []int {
	return query(tree, -1<<31, 1<<31-1)
}

func min(tree rangeTree) int {
	if tree.ltree != nil {
		return min(*tree.ltree)
	}
	return tree.val
}

func max(tree rangeTree) int {
	if tree.rtree != nil {
		return max(*tree.rtree)
	}
	return tree.val
}

func search(tree rangeTree, val int) bool {
	if tree.val > val && tree.ltree != nil {
		return search(*tree.ltree, val)
	}
	if tree.val == val {
		return true
	}
	if tree.val < val && tree.rtree != nil {
		return search(*tree.rtree, val)
	}
	return false
}

func query(tree rangeTree, l, r int) []int {
	res := []int{}
	if tree.val > l && tree.ltree != nil {
		res = append(res, query(*tree.ltree, l, r)...)
	}
	if tree.val >= l && tree.val <= r {
		res = append(res, tree.val)
	}
	if tree.val < r && tree.rtree != nil {
		res = append(res, query(*tree.rtree, l, r)...)
	}
	return res
}

func reverseQuery(tree rangeTree, l, r int) []int {
	res := []int{}
	if tree.val < r && tree.rtree != nil {
		res = append(res, reverseQuery(*tree.rtree, l, r)...)
	}
	if tree.val >= l && tree.val <= r {
		res = append(res, tree.val)
	}
	if tree.val > l && tree.ltree != nil {
		res = append(res, reverseQuery(*tree.ltree, l, r)...)
	}
	return res
}
