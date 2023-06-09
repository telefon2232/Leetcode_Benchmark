package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question144 struct {
	para144
	ans144
}

// para 是参数
// one 代表第一个参数
type para144 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans144 struct {
	one []int
}

func Benchmark_Problem144(b *testing.B) {

	qs := []question144{

		{
			para144{[]int{}},
			ans144{[]int{}},
		},

		{
			para144{[]int{1}},
			ans144{[]int{1}},
		},

		{
			para144{[]int{1, structures.NULL, 2, 3}},
			ans144{[]int{1, 2, 3}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans144, q.para144

				root := structures.Ints2TreeNode(p.one)
				(preorderTraversal(root))
			}
		}
	}
}
