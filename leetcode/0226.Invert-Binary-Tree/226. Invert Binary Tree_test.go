package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question226 struct {
	para226
	ans226
}

// para 是参数
// one 代表第一个参数
type para226 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans226 struct {
	one []int
}

func Benchmark_Problem226(b *testing.B) {

	qs := []question226{

		{
			para226{[]int{}},
			ans226{[]int{}},
		},

		{
			para226{[]int{1}},
			ans226{[]int{1}},
		},

		{
			para226{[]int{4, 2, 7, 1, 3, 6, 9}},
			ans226{[]int{4, 7, 2, 9, 6, 3, 1}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans226, q.para226

				root := structures.Ints2TreeNode(p.one)
				(structures.Tree2ints(invertTree(root)))
			}
		}
	}
}
