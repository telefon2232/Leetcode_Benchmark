package leetcode

import (
	"testing"
)

type question992 struct {
	para992
	ans992
}

// para 是参数
// one 代表第一个参数
type para992 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans992 struct {
	one int
}

func Benchmark_Problem992(b *testing.B) {

	qs := []question992{

		{
			para992{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 1},
			ans992{36},
		},

		{
			para992{[]int{2, 1, 1, 1, 2}, 1},
			ans992{8},
		},

		{
			para992{[]int{1, 2}, 1},
			ans992{2},
		},

		{
			para992{[]int{1, 2, 1, 2, 3}, 2},
			ans992{7},
		},

		{
			para992{[]int{1, 2, 1, 3, 4}, 3},
			ans992{3},
		},

		{
			para992{[]int{1}, 5},
			ans992{1},
		},

		{
			para992{[]int{}, 10},
			ans992{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans992, q.para992
				(subarraysWithKDistinct(p.one, p.k))
			}
		}
	}
}
