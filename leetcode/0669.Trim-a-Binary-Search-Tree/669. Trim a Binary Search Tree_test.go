package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question669 struct {
	para669
	ans669
}

// para 是参数
// one 代表第一个参数
type para669 struct {
	one  []int
	low  int
	high int
}

// ans 是答案
// one 代表第一个答案
type ans669 struct {
	one []int
}

func Benchmark_Problem669(b *testing.B) {

	qs := []question669{

		{
			para669{[]int{1, 0, 2}, 1, 2},
			ans669{[]int{1, structures.NULL, 2}},
		},

		{
			para669{[]int{3, 0, 4, structures.NULL, 2, structures.NULL, structures.NULL, 1}, 1, 3},
			ans669{[]int{3, 2, structures.NULL, 1}},
		},

		{
			para669{[]int{1}, 1, 2},
			ans669{[]int{1}},
		},

		{
			para669{[]int{1, structures.NULL, 2}, 1, 3},
			ans669{[]int{1, structures.NULL, 2}},
		},

		{
			para669{[]int{1, structures.NULL, 2}, 2, 4},
			ans669{[]int{2}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans669, q.para669

				root := structures.Ints2TreeNode(p.one)
				(structures.Tree2ints(trimBST(root, p.low, p.high)))
			}
		}
	}
}
