package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question82 struct {
	para82
	ans82
}

// para 是参数
// one 代表第一个参数
type para82 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans82 struct {
	one []int
}

func Benchmark_Problem82(b *testing.B) {

	qs := []question82{

		{
			para82{[]int{1, 1, 2, 2, 3, 4, 4, 4}},
			ans82{[]int{3}},
		},

		{
			para82{[]int{1, 1, 1, 1, 1, 1}},
			ans82{[]int{}},
		},

		{
			para82{[]int{1, 1, 1, 2, 3}},
			ans82{[]int{2, 3}},
		},

		{
			para82{[]int{1}},
			ans82{[]int{1}},
		},

		{
			para82{[]int{}},
			ans82{[]int{}},
		},

		{
			para82{[]int{1, 2, 2, 2, 2}},
			ans82{[]int{1}},
		},

		{
			para82{[]int{1, 1, 2, 3, 3, 4, 5, 5, 6}},
			ans82{[]int{2, 4, 6}},
		},

		{
			para82{[]int{1, 1, 2, 3, 3, 4, 5, 6}},
			ans82{[]int{2, 4, 5, 6}},
		},

		{
			para82{[]int{0, 1, 2, 2, 3, 4}},
			ans82{[]int{0, 1, 2, 2, 3, 4}},
		},
	}

	for _, q := range qs {
		_, p := q.ans82, q.para82
		(structures.List2Ints(deleteDuplicates1(structures.Ints2List(p.one))))
	}
}
