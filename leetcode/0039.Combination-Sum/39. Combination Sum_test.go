package leetcode

import (
	"testing"
)

type question39 struct {
	para39
	ans39
}

// para 是参数
// one 代表第一个参数
type para39 struct {
	n []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans39 struct {
	one [][]int
}

func Benchmark_Problem39(b *testing.B) {

	qs := []question39{

		{
			para39{[]int{2, 3, 6, 7}, 7},
			ans39{[][]int{{7}, {2, 2, 3}}},
		},

		{
			para39{[]int{2, 3, 5}, 8},
			ans39{[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans39, q.para39
				(combinationSum(p.n, p.k))
			}
		}
	}
}
