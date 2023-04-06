package leetcode

import (
	"testing"
)

type question40 struct {
	para40
	ans40
}

// para 是参数
// one 代表第一个参数
type para40 struct {
	n []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans40 struct {
	one [][]int
}

func Benchmark_Problem40(b *testing.B) {

	qs := []question40{

		{
			para40{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
			ans40{[][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}}},
		},

		{
			para40{[]int{2, 5, 2, 1, 2}, 5},
			ans40{[][]int{{1, 2, 2}, {5}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans40, q.para40
				(combinationSum2(p.n, p.k))
			}
		}
	}
}
