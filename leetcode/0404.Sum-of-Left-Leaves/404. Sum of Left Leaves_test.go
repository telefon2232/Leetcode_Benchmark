package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question404 struct {
	para404
	ans404
}

// para 是参数
// one 代表第一个参数
type para404 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans404 struct {
	one int
}

func Benchmark_Problem404(b *testing.B) {

	qs := []question404{

		{
			para404{[]int{}},
			ans404{0},
		},

		{
			para404{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans404{24},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans404, q.para404

				root := structures.Ints2TreeNode(p.one)
				(sumOfLeftLeaves(root))
			}
		}
	}
}
