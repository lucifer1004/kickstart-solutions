package rangetree1d

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRangeTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Range Tree 1D Test Suite")
}

var _ = Describe("Range Tree", func() {
	var (
		tree RangeTree
		vals []int
	)

	BeforeEach(func() {
		vals = []int{7, 1, 9, -2, 10, 19, 8, 2, 3, 4}
		tree = Build(vals)
	})

	Describe("Build", func() {
		It("should panic when vals is empty", func() {
			Expect(func() { Build([]int{}) }).To(Panic())
		})
	})

	Describe("Traverse", func() {
		It("should return [-2, 1, 2, 3, 4, 7, 8, 9, 10, 19]", func() {
			Expect(Traverse(tree)).To(Equal([]int{-2, 1, 2, 3, 4, 7, 8, 9, 10, 19}))
		})
	})

	Describe("Insert", func() {
		It("12 should be handled correctly", func() {
			Insert(&tree, 12)
			Expect(Traverse(tree)).To(Equal([]int{-2, 1, 2, 3, 4, 7, 8, 9, 10, 12, 19}))
		})
	})

	Describe("Delete", func() {
		It("2 should be handled correctly", func() {
			Delete(&tree, 2)
			Expect(Traverse(tree)).To(Equal([]int{-2, 1, 3, 4, 7, 8, 9, 10, 19}))
		})
		It("19 should be handled correctly", func() {
			Delete(&tree, 19)
			Expect(Traverse(tree)).To(Equal([]int{-2, 1, 2, 3, 4, 7, 8, 9, 10}))
		})
	})

	Describe("Min", func() {
		It("should return -2", func() {
			Expect(Min(tree)).To(Equal(-2))
		})
	})

	Describe("Max", func() {
		It("should return 19", func() {
			Expect(Max(tree)).To(Equal(19))
		})
	})

	Describe("Search", func() {
		It("12 should return false", func() {
			Expect(Search(tree, 12)).To(Equal(false))
		})
		It("19 should return true", func() {
			Expect(Search(tree, 19)).To(Equal(true))
		})
	})

	Describe("Query", func() {
		It("[0, 5] should return [1, 2, 3, 4]", func() {
			Expect(Query(tree, 0, 5)).To(Equal([]int{1, 2, 3, 4}))
		})

		It("[-100, -1] should return [-2]", func() {
			Expect(Query(tree, -100, -1)).To(Equal([]int{-2}))
		})
	})

	Describe("Reverse Query", func() {
		It("[9, 15] should return [10, 9]", func() {
			Expect(ReverseQuery(tree, 9, 15)).To(Equal([]int{10, 9}))
		})

		It("[7, 100] should return [19, 10, 9, 8, 7]", func() {
			Expect(ReverseQuery(tree, 7, 100)).To(Equal([]int{19, 10, 9, 8, 7}))
		})
	})
})
