package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question102 struct {
	para102
	ans102
}

// para 是参数
// one 代表第一个参数
type para102 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans102 struct {
	one [][]int
}

func Benchmark_Problem102(b *testing.B) {

	qs := []question102{

		{
			para102{[]int{}},
			ans102{[][]int{}},
		},

		{
			para102{[]int{1}},
			ans102{[][]int{{1}}},
		},

		{
			para102{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans102{[][]int{{3}, {9, 20}, {15, 7}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans102, q.para102

				root := structures.Ints2TreeNode(p.one)
				(levelOrder(root))
			}
		}
	}
}
