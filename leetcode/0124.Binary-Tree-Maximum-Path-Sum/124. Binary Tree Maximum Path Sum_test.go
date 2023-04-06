package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question124 struct {
	para124
	ans124
}

// para 是参数
// one 代表第一个参数
type para124 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans124 struct {
	one int
}

func Benchmark_Problem124(b *testing.B) {

	qs := []question124{

		{
			para124{[]int{}},
			ans124{0},
		},

		{
			para124{[]int{1}},
			ans124{1},
		},

		{
			para124{[]int{1, 2, 3}},
			ans124{6},
		},

		{
			para124{[]int{-10, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans124{42},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans124, q.para124

				root := structures.Ints2TreeNode(p.one)
				(maxPathSum(root))
			}
		}
	}
}
