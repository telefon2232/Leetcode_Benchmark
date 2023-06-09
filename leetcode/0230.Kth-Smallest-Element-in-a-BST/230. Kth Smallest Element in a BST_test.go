package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question230 struct {
	para230
	ans230
}

// para 是参数
// one 代表第一个参数
type para230 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans230 struct {
	one int
}

func Benchmark_Problem230(b *testing.B) {

	qs := []question230{

		{
			para230{[]int{}, 0},
			ans230{0},
		},

		{
			para230{[]int{3, 1, 4, structures.NULL, 2}, 1},
			ans230{1},
		},

		{
			para230{[]int{5, 3, 6, 2, 4, structures.NULL, structures.NULL, 1}, 3},
			ans230{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans230, q.para230

				root := structures.Ints2TreeNode(p.one)
				(kthSmallest(root, p.k))
			}
		}
	}
}
