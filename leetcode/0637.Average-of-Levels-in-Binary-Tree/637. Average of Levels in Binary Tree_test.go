package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question637 struct {
	para637
	ans637
}

// para 是参数
// one 代表第一个参数
type para637 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans637 struct {
	one [][]int
}

func Benchmark_Problem637(b *testing.B) {

	qs := []question637{

		{
			para637{[]int{}},
			ans637{[][]int{}},
		},

		{
			para637{[]int{1}},
			ans637{[][]int{{1}}},
		},

		{
			para637{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans637{[][]int{{3}, {9, 20}, {15, 7}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans637, q.para637

				root := structures.Ints2TreeNode(p.one)
				(averageOfLevels(root))
			}
		}
	}
}
