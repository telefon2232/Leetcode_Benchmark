package leetcode

import (
	"testing"
)

type question174 struct {
	para174
	ans174
}

// para 是参数
// one 代表第一个参数
type para174 struct {
	s [][]int
}

// ans 是答案
// one 代表第一个答案
type ans174 struct {
	one int
}

func Benchmark_Problem174(b *testing.B) {

	qs := []question174{

		{
			para174{[][]int{{2, 1}, {1, -1}}},
			ans174{1},
		},

		{
			para174{[][]int{{-3, 5}}},
			ans174{4},
		},

		{
			para174{[][]int{{100}}},
			ans174{1},
		},

		{
			para174{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}},
			ans174{7},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans174, q.para174
				(calculateMinimumHP1(p.s))
			}
		}
	}
}
