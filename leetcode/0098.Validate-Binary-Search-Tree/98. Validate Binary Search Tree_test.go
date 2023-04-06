package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question98 struct {
	para98
	ans98
}

// para 是参数
// one 代表第一个参数
type para98 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans98 struct {
	one bool
}

func Benchmark_Problem98(b *testing.B) {

	qs := []question98{

		{
			para98{[]int{10, 5, 15, structures.NULL, structures.NULL, 6, 20}},
			ans98{false},
		},

		{
			para98{[]int{}},
			ans98{true},
		},

		{
			para98{[]int{2, 1, 3}},
			ans98{true},
		},

		{
			para98{[]int{5, 1, 4, structures.NULL, structures.NULL, 3, 6}},
			ans98{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans98, q.para98

				rootOne := structures.Ints2TreeNode(p.one)
				(isValidBST(rootOne))
			}
		}
	}
}
