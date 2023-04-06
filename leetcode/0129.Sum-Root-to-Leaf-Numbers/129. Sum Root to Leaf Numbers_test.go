package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question129 struct {
	para129
	ans129
}

// para 是参数
// one 代表第一个参数
type para129 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans129 struct {
	one int
}

func Benchmark_Problem129(b *testing.B) {

	qs := []question129{

		{
			para129{[]int{}},
			ans129{0},
		},

		{
			para129{[]int{1, 2, 3}},
			ans129{25},
		},

		{
			para129{[]int{4, 9, 0, 5, 1}},
			ans129{1026},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans129, q.para129

				root := structures.Ints2TreeNode(p.one)
				(sumNumbers(root))
			}
		}
	}
}
