package rangetree1d

// RangeTree ...
type RangeTree struct {
	val          int
	ltree, rtree *RangeTree
}

// Build ...
func Build(vals []int) RangeTree {
	if len(vals) == 0 {
		panic("Cannot initialize an empty range tree")
	}
	tree := RangeTree{val: vals[0]}
	for i := 1; i < len(vals); i++ {
		Insert(&tree, vals[i])
	}
	return tree
}

// Insert ...
func Insert(tree *RangeTree, val int) {
	if val == tree.val {
		return
	}
	if val < tree.val {
		if tree.ltree != nil {
			Insert(tree.ltree, val)
		} else {
			tree.ltree = &RangeTree{
				val: val,
			}
		}
	} else {
		if tree.rtree != nil {
			Insert(tree.rtree, val)
		} else {
			tree.rtree = &RangeTree{
				val: val,
			}
		}
	}
}

// Delete ...
func Delete(tree *RangeTree, val int) *RangeTree {
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
			tree.ltree = Delete(tree.ltree, val)
		}
	} else {
		if tree.rtree != nil {
			tree.rtree = Delete(tree.rtree, val)
		}
	}
	return tree
}

// Traverse ...
func Traverse(tree RangeTree) []int {
	return Query(tree, -1<<31, 1<<31-1)
}

// Min ...
func Min(tree RangeTree) int {
	if tree.ltree != nil {
		return Min(*tree.ltree)
	}
	return tree.val
}

// Max ...
func Max(tree RangeTree) int {
	if tree.rtree != nil {
		return Max(*tree.rtree)
	}
	return tree.val
}

// Search ...
func Search(tree RangeTree, val int) bool {
	if tree.val > val && tree.ltree != nil {
		return Search(*tree.ltree, val)
	}
	if tree.val == val {
		return true
	}
	if tree.val < val && tree.rtree != nil {
		return Search(*tree.rtree, val)
	}
	return false
}

// Query ...
func Query(tree RangeTree, l, r int) []int {
	res := []int{}
	if tree.val > l && tree.ltree != nil {
		res = append(res, Query(*tree.ltree, l, r)...)
	}
	if tree.val >= l && tree.val <= r {
		res = append(res, tree.val)
	}
	if tree.val < r && tree.rtree != nil {
		res = append(res, Query(*tree.rtree, l, r)...)
	}
	return res
}

// ReverseQuery ...
func ReverseQuery(tree RangeTree, l, r int) []int {
	res := []int{}
	if tree.val < r && tree.rtree != nil {
		res = append(res, ReverseQuery(*tree.rtree, l, r)...)
	}
	if tree.val >= l && tree.val <= r {
		res = append(res, tree.val)
	}
	if tree.val > l && tree.ltree != nil {
		res = append(res, ReverseQuery(*tree.ltree, l, r)...)
	}
	return res
}
